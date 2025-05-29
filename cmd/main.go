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
	mode     string
}

func initialModel() model {
	repo, err := domain.NewSQLiteRepository("zettl.db")
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
		case "enter", " ":
			// Save a new note to the database when enter/space is pressed
		Width(22)

	header := style.Render("Zettl Notes")

	s := header + "\n\n"
	s += "What would you like to do?\n\n"

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
