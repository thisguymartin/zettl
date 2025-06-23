package internal

import (
	"github.com/charmbracelet/log"
	"time"
)

type WindowType int

const (
	MainMenuWindow WindowType = iota
	NoteListWindow
	NoteEditWindow
	SearchWindow
)

type Note struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Tags      string    `db:"tags"`
}

type App struct {
	Name           string
	currentWindow  WindowType
	notes          []Note
	filteredNotes  []Note
	cursor         int
	searchQuery    string
	noteContent    string
	noteTitle      string
	selectedNote   Note
	noteRepository NoteRepository
	width          int
	height         int
}

func New(repo NoteRepository) (*App, error) {
	notes, err := repo.GetAll()
	if err != nil {
		log.Error("failed to get notes: %v", "error", err)
		notes = []Note{}
	}

	return &App{
		Name:           "zettl",
		currentWindow:  MainMenuWindow,
		notes:          notes,
		filteredNotes:  notes,
		noteRepository: repo,
		width:          80,
		height:         24,
	}, nil
}
