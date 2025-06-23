package database

import (
	"database/sql"
	"fmt"
	"time"

	internal "thisguymartin/zettl/internal"

	_ "github.com/tursodatabase/go-libsql"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(dbPath string) (*SQLiteRepository, error) {
	db, err := sql.Open("libsql", "file:"+dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	repo := &SQLiteRepository{db: db}
	if err := repo.createTable(); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return repo, nil
}

func (r *SQLiteRepository) createTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		tags TEXT DEFAULT ''
	);
	`
	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) Create(note internal.Note) error {
	query := `
	INSERT INTO notes (title, content, tags, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`
	now := time.Now()
	result, err := r.db.Exec(query, note.Title, note.Content, note.Tags, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	note.ID = int(id)
	note.CreatedAt = now
	note.UpdatedAt = now
	return nil
}

func (r *SQLiteRepository) GetAll() ([]internal.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at, tags FROM notes ORDER BY updated_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []internal.Note
	for rows.Next() {
		var note internal.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func (r *SQLiteRepository) GetByID(id int) (internal.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at, tags FROM notes WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var note internal.Note
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags)
	if err != nil {
		return internal.Note{}, err
	}

	return note, nil
}

func (r *SQLiteRepository) Update(note internal.Note) error {
	query := `
	UPDATE notes 
	SET title = ?, content = ?, tags = ?, updated_at = ?
	WHERE id = ?
	`
	note.UpdatedAt = time.Now()
	_, err := r.db.Exec(query, note.Title, note.Content, note.Tags, note.UpdatedAt, note.ID)
	return err
}

func (r *SQLiteRepository) Delete(id int) error {
	query := `DELETE FROM notes WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SQLiteRepository) Search(query string) ([]internal.Note, error) {
	searchQuery := `
	SELECT id, title, content, created_at, updated_at, tags 
	FROM notes 
	WHERE title LIKE ? OR content LIKE ? OR tags LIKE ?
	ORDER BY updated_at DESC
	`
	pattern := "%" + query + "%"
	rows, err := r.db.Query(searchQuery, pattern, pattern, pattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []internal.Note
	for rows.Next() {
		var note internal.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func (r *SQLiteRepository) Close() error {
	return r.db.Close()
}
