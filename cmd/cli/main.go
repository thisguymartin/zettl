package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	domain "thisguymartin/zettl/internal/domain"
)

type model struct {
	notes    []string
	cursor   int
	selected map[int]struct{}
	mode     string // "view" or "input"
	input    string // buffer for new note
}

func initialModel() model {
	repo, err := domain.NewSQLiteRepository("libsql://zettl.db")
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	notes, err := repo.GetAll()
	if err != nil {
		log.Printf("failed to load notes: %v", err)
	}

	noteTitles := make([]string, len(notes))
	for i, n := range notes {
		noteTitles[i] = n.Title
	}

	return model{
		notes:    noteTitles,
		selected: make(map[int]struct{}),
		mode:     "view",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.mode {
		case "view":
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.notes)-1 {
					m.cursor++
				}
			case "n":
				m.mode = "input"
				m.input = ""
				return m, nil
			case "enter", " ":
				// Save a new note to the database when enter/space is pressed
				repo, err := domain.NewSQLiteRepository("libsql://zettl.db")
				if err != nil {
					log.Printf("failed to open db: %v", err)
					return m, nil
				}
				note := &domain.Note{
					Title:   m.notes[m.cursor],
					Content: m.notes[m.cursor],
					Tags:    "cli",
				}
				err = repo.Create(note)
				if err != nil {
					log.Printf("failed to save note: %v", err)
				} else {
					log.Printf("note saved: %s", note.Title)
				}
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}
			}
		case "input":
			switch msg.Type {
			case tea.KeyEnter:
				if m.input != "" {
					repo, err := domain.NewSQLiteRepository("libsql://zettl.db")
					if err == nil {
						note := &domain.Note{Title: m.input, Content: m.input, Tags: "cli"}
						repo.Create(note)
						notes, _ := repo.GetAll()
						noteTitles := make([]string, len(notes))
						for i, n := range notes {
							noteTitles[i] = n.Title
						}
						m.notes = noteTitles
					}
					m.input = ""
				}
				m.mode = "view"
				return m, nil
			case tea.KeyEsc:
				m.input = ""
				m.mode = "view"
				return m, nil
			case tea.KeyBackspace, tea.KeyCtrlH:
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
				return m, nil
			default:
				if msg.Type == tea.KeyRunes {
					m.input += msg.String()
				}
				return m, nil
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)

	header := style.Render("Zettl Notes")

	s := header + "\n\n"
	if m.mode == "input" {
		s += "Type your note and press Enter to save. (Esc to cancel)\n\n"
		s += "> " + m.input + "\n"
		return s
	}

	s += "What would you like to do? (n: new note)\n\n"

	for i, note := range m.notes {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, note)
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
