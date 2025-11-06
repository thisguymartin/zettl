package domain

import (
	"time"
)

type Note struct {
	ID        int       `db:"id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Tags      string    `db:"tags"`
	Embedding string    `db:"embedding"` // JSON-encoded vector
}

type ChatMessage struct {
	ID        int       `db:"id"`
	Role      string    `db:"role"` // "user" or "assistant"
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type NoteRepository interface {
	Create(note *Note) error
	GetAll() ([]Note, error)
	GetByID(id int) (*Note, error)
	Update(note *Note) error
	Delete(id int) error
	Search(query string) ([]Note, error)
	UpdateEmbedding(noteID int, embedding string) error
	GetNotesWithEmbeddings() ([]Note, error)
	SaveChatMessage(role, content string) error
	GetChatHistory(limit int) ([]ChatMessage, error)
	ClearChatHistory() error
}
