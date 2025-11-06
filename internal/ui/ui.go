package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type WindowType int

const (
	MainMenuWindow WindowType = iota
	NoteListWindow
	NoteEditWindow
	SearchWindow
	ChatWindow
	AIMenuWindow
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
	// AI-related fields
	chatInput    string
	chatHistory  []ChatMessage
	isProcessing bool
	errorMsg     string
	aiService    interface{} // Will hold *ai.AIService
}

func NewUIModel(repo NoteRepository) (*UIModel, error) {
	notes, err := repo.GetAll()
	if err != nil {
		log.Error("failed to get notes: %v", "error", err)
		notes = []Note{}
	}

	return &UIModel{
		currentWindow: MainMenuWindow,
		notes:         notes,
		filteredNotes: notes,
		repo:          repo,
		width:         80,
		height:        24,
		chatHistory:   []ChatMessage{},
	}, nil
}

// SetAIService sets the AI service for the model
func (m *UIModel) SetAIService(service interface{}) {
	m.aiService = service
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
	case ChatWindow:
		return m.viewChat()
	case AIMenuWindow:
		return m.viewAIMenu()
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
		case ChatWindow:
			return m.updateChat(msg)
		case AIMenuWindow:
			return m.updateAIMenu(msg)
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
	case "4", "a":
		m.currentWindow = AIMenuWindow
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

	headerStyle := lipgloss.NewStyle().AlignVertical(lipgloss.Center).
		Width(m.width).
		Align(lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(1, 2).
		MarginBottom(1).
		Width(m.width).
		Render(`
╔══════════════════════════════════════════════════════════════╗
║  			███████╗███████╗████████╗████████╗██╗         	  ║
║  			╚══███╔╝██╔════╝╚══██╔══╝╚══██╔══╝██║         	  ║
║  			  ███╔╝ █████╗     ██║      ██║   ██║         	  ║
║  			 ███╔╝  ██╔══╝     ██║      ██║   ██║         	  ║
║  			███████╗███████╗   ██║      ██║   ███████╗    	  ║
║  			╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝    	  ║
╚══════════════════════════════════════════════════════════════╝`)

	menuStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		MarginTop(1).
		Align(lipgloss.Center)

	menu := menuStyle.Render(`Welcome to your digital notebook!

1. New Note (n)     - Create a new note
2. List Notes (l)   - Browse existing notes
3. Search (s)       - Find notes with fuzzy search
4. AI Features (a)  - Chat, embeddings, and semantic search

Press the number or letter key to navigate.
Press 'q' to quit.`)

	return lipgloss.JoinVertical(lipgloss.Left, headerStyle, menu)
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

func (m UIModel) viewAIMenu() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2).
		MarginBottom(1)

	menuStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00D9FF")).
		Padding(1, 2).
		MarginTop(1).
		Align(lipgloss.Center)

	header := headerStyle.Render("🤖 AI Features")

	menu := menuStyle.Render(`AI-Powered Features:

1. Chat (c)              - Chat with AI about your notes
2. Generate Embeddings   - Create embeddings for semantic search
3. Semantic Search       - Find similar notes using AI
4. Clear Chat History    - Reset conversation

Press the number or letter key to navigate.
Press 'Esc' to go back.`)

	if m.errorMsg != "" {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Padding(1, 2)
		error := errorStyle.Render("⚠️  " + m.errorMsg)
		return lipgloss.JoinVertical(lipgloss.Left, header, menu, error)
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, menu)
}

func (m UIModel) updateAIMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = MainMenuWindow
		m.errorMsg = ""
	case "1", "c":
		// Load chat history
		if history, err := m.repo.GetChatHistory(50); err == nil {
			m.chatHistory = history
		}
		m.currentWindow = ChatWindow
		m.chatInput = ""
		m.errorMsg = ""
	case "2":
		// Generate embeddings - to be implemented
		m.errorMsg = "Embedding generation - coming soon!"
	case "3":
		// Semantic search - to be implemented
		m.errorMsg = "Semantic search - coming soon!"
	case "4":
		// Clear chat history
		if err := m.repo.ClearChatHistory(); err == nil {
			m.chatHistory = []ChatMessage{}
			m.errorMsg = "Chat history cleared!"
		} else {
			m.errorMsg = "Failed to clear chat history"
		}
	}
	return m, nil
}

func (m UIModel) viewChat() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 2).
		MarginBottom(1)

	chatStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00D9FF")).
		Padding(1, 2).
		Width(m.width - 4).
		Height(m.height - 12)

	inputStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#32CD32")).
		Padding(1, 2).
		Width(m.width - 4)

	header := headerStyle.Render("💬 AI Chat")

	// Render chat history
	var chatContent strings.Builder
	if len(m.chatHistory) == 0 {
		chatContent.WriteString("No messages yet. Start a conversation!\n\n")
		chatContent.WriteString("Tip: The AI has access to your notes for context.")
	} else {
		for _, msg := range m.chatHistory {
			role := "You"
			if msg.Role == "assistant" {
				role = "AI"
			}
			chatContent.WriteString(fmt.Sprintf("[%s]: %s\n\n", role, msg.Content))
		}
	}

	if m.isProcessing {
		chatContent.WriteString("[AI is thinking...]\n")
	}

	chat := chatStyle.Render(chatContent.String())

	// Render input box
	input := inputStyle.Render(fmt.Sprintf("Message: %s█", m.chatInput))

	footer := "\nType your message • Enter: Send • Esc: Back"

	if m.errorMsg != "" {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Padding(1, 0)
		error := errorStyle.Render("⚠️  " + m.errorMsg)
		return lipgloss.JoinVertical(lipgloss.Left, header, chat, input, error, footer)
	}

	return lipgloss.JoinVertical(lipgloss.Left, header, chat, input, footer)
}

func (m UIModel) updateChat(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = AIMenuWindow
		m.chatInput = ""
		m.errorMsg = ""
	case "enter":
		if m.chatInput != "" && !m.isProcessing {
			// Save user message
			userMsg := m.chatInput
			if err := m.repo.SaveChatMessage("user", userMsg); err != nil {
				m.errorMsg = "Failed to save message"
				return m, nil
			}

			// Add to history
			m.chatHistory = append(m.chatHistory, ChatMessage{
				Role:    "user",
				Content: userMsg,
			})

			m.chatInput = ""
			m.isProcessing = true
			m.errorMsg = ""

			// TODO: In a real implementation, this would call the AI service asynchronously
			// For now, just add a placeholder response
			response := "AI integration requires OpenAI API key. Please set OPENAI_API_KEY environment variable or configure it in ~/.zettl/config.json"

			if err := m.repo.SaveChatMessage("assistant", response); err != nil {
				m.errorMsg = "Failed to save AI response"
			} else {
				m.chatHistory = append(m.chatHistory, ChatMessage{
					Role:    "assistant",
					Content: response,
				})
			}
			m.isProcessing = false
		}
	case "backspace", "ctrl+h":
		if len(m.chatInput) > 0 {
			m.chatInput = m.chatInput[:len(m.chatInput)-1]
		}
	default:
		if msg.Type == tea.KeyRunes && !m.isProcessing {
			m.chatInput += msg.String()
		}
	}
	return m, nil
}
