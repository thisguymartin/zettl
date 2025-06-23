package internal

type NoteRepository interface {
	Create(note Note) error
	GetAll() ([]Note, error)
	GetByID(id int) (Note, error)
	Update(note Note) error
	Delete(id int) error
	Search(query string) ([]Note, error)
}
