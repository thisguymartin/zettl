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
}

type NoteRepository interface {
	Create(note *Note) error
	GetAll() ([]Note, error)
	GetByID(id int) (*Note, error)
	Update(note *Note) error
	Delete(id int) error
	Search(query string) ([]Note, error)
}
