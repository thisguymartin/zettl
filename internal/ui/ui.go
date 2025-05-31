package domain

import (
	"fmt"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WindowType int

const (
	MainMenuWindow WindowType = iota
	NoteListWindow
	NoteEditWindow
	SearchWindow
)

type UIModel struct {
	currentWindow WindowType
	notes         []Note
	filteredNotes []Note
	cursor        int
	searchQuery   string
	noteContent   string
	noteTitle     string
	selectedNote  *Note
	repo          NoteRepository
	width         int
	height        int
}

func NewUIModel(repo NoteRepository) (*UIModel, error) {
	notes, err := repo.GetAll()
	if err != nil {
		log.Printf("failed to load notes: %v", err)
		notes = []Note{}
	}

	return &UIModel{
		currentWindow: MainMenuWindow,
		notes:         notes,
		filteredNotes: notes,
		repo:          repo,
		width:         80,
		height:        24,
	}, nil
}

func (m UIModel) Init() tea.Cmd {
	return nil
}

func (m UIModel) View() string {
	switch m.currentWindow {
	case MainMenuWindow:
		return m.viewMainMenu()
	case NoteListWindow:
		return m.viewNoteList()
	case NoteEditWindow:
		return m.viewNoteEdit()
	case SearchWindow:
		return m.viewSearch()
	}
	return ""
}

func (m UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch m.currentWindow {
		case MainMenuWindow:
			return m.updateMainMenu(msg)
		case NoteListWindow:
			return m.updateNoteList(msg)
		case NoteEditWindow:
			return m.updateNoteEdit(msg)
		case SearchWindow:
			return m.updateSearch(msg)
		}
	}
	return m, nil
}

func (m *UIModel) refreshNotes() {
	notes, err := m.repo.GetAll()
	if err != nil {
		log.Printf("failed to refresh notes: %v", err)
		return
	}
	m.notes = notes
	m.applyFilter()
}

func (m *UIModel) applyFilter() {
	if m.searchQuery == "" {
		m.filteredNotes = m.notes
		return
	}

	filtered := []Note{}
	query := strings.ToLower(m.searchQuery)
	for _, note := range m.notes {
		if strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Content), query) ||
			strings.Contains(strings.ToLower(note.Tags), query) {
			filtered = append(filtered, note)
		}
	}
	m.filteredNotes = filtered
}

func (m UIModel) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "1", "n":
		m.currentWindow = NoteEditWindow
		m.noteTitle = ""
		m.noteContent = ""
		m.selectedNote = nil
	case "2", "l":
		m.currentWindow = NoteListWindow
		m.cursor = 0
	case "3", "s":
		m.currentWindow = SearchWindow
		m.searchQuery = ""
		m.cursor = 0
	}
	return m, nil
}

func (m UIModel) updateNoteList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		m.currentWindow = MainMenuWindow
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.filteredNotes)-1 {
			m.cursor++
		}
	case "enter":
		if len(m.filteredNotes) > 0 && m.cursor < len(m.filteredNotes) {
			m.selectedNote = &m.filteredNotes[m.cursor]
			m.noteTitle = m.selectedNote.Title
			m.noteContent = m.selectedNote.Content
			m.currentWindow = NoteEditWindow
		}
	case "/":
		m.currentWindow = SearchWindow
		m.searchQuery = ""
	}
	return m, nil
}

func (m UIModel) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		m.currentWindow = NoteListWindow
		m.searchQuery = ""
		m.applyFilter()
	case "enter":
		m.currentWindow = NoteListWindow
		m.cursor = 0
		m.applyFilter()
	case "backspace", "ctrl+h":
		if len(m.searchQuery) > 0 {
			m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
			m.applyFilter()
		}
	default:
		if msg.Type == tea.KeyRunes {
			m.searchQuery += msg.String()
			m.applyFilter()
		}
	}
	return m, nil
}

func (m UIModel) updateNoteEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = NoteListWindow
	case "ctrl+s":
		return m.saveNote()
	case "backspace", "ctrl+h":
		if len(m.noteContent) > 0 {
			m.noteContent = m.noteContent[:len(m.noteContent)-1]
		}
	case "enter":
		m.noteContent += "\n"
	case "tab":
		m.noteContent += "    "
	case " ":
		m.noteContent += " "
	default:
		if msg.Type == tea.KeyRunes {
			m.noteContent += msg.String()
		}
	}
	return m, nil
}

func (m UIModel) saveNote() (tea.Model, tea.Cmd) {
	if m.noteContent == "" {
		return m, nil
	}

	if m.noteTitle == "" {
		lines := strings.Split(m.noteContent, "\n")
		if len(lines) > 0 && strings.TrimSpace(lines[0]) != "" {
			m.noteTitle = strings.TrimSpace(lines[0])
			if len(m.noteTitle) > 50 {
				m.noteTitle = m.noteTitle[:50] + "..."
			}
		} else {
			m.noteTitle = "Untitled Note"
		}
	}

	if m.selectedNote != nil {
		m.selectedNote.Title = m.noteTitle
		m.selectedNote.Content = m.noteContent
		m.selectedNote.UpdatedAt = time.Now()
		err := m.repo.Update(m.selectedNote)
		if err != nil {
			log.Printf("failed to update note: %v", err)
		}
	} else {
		note := &Note{
			Title:     m.noteTitle,
			Content:   m.noteContent,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Tags:      "notebook",
		}
		err := m.repo.Create(note)
		if err != nil {
			log.Printf("failed to create note: %v", err)
		}
	}

	m.refreshNotes()
	m.currentWindow = NoteListWindow
	return m, nil
}

func (m UIModel) viewMainMenu() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		MarginBottom(1)

	menuStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		MarginTop(1)

	header := headerStyle.Render("📝 Zettl Notebook")

	menu := menuStyle.Render(`Welcome to your digital notebook!

1. New Note (n)     - Create a new note
2. List Notes (l)   - Browse existing notes  
3. Search (s)       - Find notes with fuzzy search

Press the number or letter key to navigate.
Press 'q' to quit.`)

	return lipgloss.JoinVertical(lipgloss.Left, header, menu)
}

func (m UIModel) viewNoteList() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2).
		MarginBottom(1)

	noteStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		Width(m.width - 4)

	selectedStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF6B6B")).
		Background(lipgloss.Color("#2A2A2A")).
		Padding(1, 2).
		Width(m.width - 4)

	header := headerStyle.Render(fmt.Sprintf("📚 Notes (%d)", len(m.filteredNotes)))

	if len(m.filteredNotes) == 0 {
		empty := noteStyle.Render("No notes found. Press 'Esc' to go back or '/' to search.")
		return lipgloss.JoinVertical(lipgloss.Left, header, empty)
	}

	var notes []string
	for i, note := range m.filteredNotes {
		dateStr := note.CreatedAt.Format("Jan 2, 2006 15:04")
		preview := strings.ReplaceAll(note.Content, "\n", " ")
		if len(preview) > 60 {
			preview = preview[:60] + "..."
		}

		content := fmt.Sprintf("📄 %s\n🕒 %s\n💭 %s", note.Title, dateStr, preview)

		if i == m.cursor {
			notes = append(notes, selectedStyle.Render(content))
		} else {
			notes = append(notes, noteStyle.Render(content))
		}
	}

	footer := "\n↑/↓ Navigate • Enter: Open • /: Search • Esc: Back"

	return lipgloss.JoinVertical(lipgloss.Left, header, strings.Join(notes, "\n"), footer)
}

func (m UIModel) viewSearch() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2).
		MarginBottom(1)

	searchStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFD700")).
		Padding(1, 2).
		Width(m.width - 4)

	header := headerStyle.Render("🔍 Search Notes")
	searchBox := searchStyle.Render(fmt.Sprintf("Search: %s|", m.searchQuery))

	results := fmt.Sprintf("Found %d notes", len(m.filteredNotes))
	footer := "\nType to search • Enter: View results • Esc: Back"

	return lipgloss.JoinVertical(lipgloss.Left, header, searchBox, results, footer)
}

func (m UIModel) viewNoteEdit() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2).
		MarginBottom(1)

	editorStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#32CD32")).
		Padding(1, 2).
		Width(m.width - 4).
		Height(m.height - 8)

	var title string
	if m.selectedNote != nil {
		title = fmt.Sprintf("✏️  Editing: %s", m.selectedNote.Title)
	} else {
		title = "✏️  New Note"
	}

	header := headerStyle.Render(title)

	content := m.noteContent + "█"
	if len(m.noteContent) == 0 {
		content = "Start typing your note...\n\n█"
	}

	editor := editorStyle.Render(content)
	footer := "\nCtrl+S: Save • Esc: Back • Enter: New line • Tab: Indent"

	return lipgloss.JoinVertical(lipgloss.Left, header, editor, footer)
}
