package database

import (
	"database/sql"
	"fmt"
	"time"

	domain "thisguymartin/zettl/internal/ui"

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
	queries := []string{
		// Notes table with embedding support
		`CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			tags TEXT DEFAULT '',
			embedding TEXT DEFAULT NULL
		);`,
		// Chat history table
		`CREATE TABLE IF NOT EXISTS chat_history (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			role TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
		// Create index for faster searches
		`CREATE INDEX IF NOT EXISTS idx_notes_updated_at ON notes(updated_at DESC);`,
	}

	for _, query := range queries {
		if _, err := r.db.Exec(query); err != nil {
			return fmt.Errorf("failed to execute migration: %w", err)
		}
	}
	return nil
}

func (r *SQLiteRepository) Create(note *domain.Note) error {
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

func (r *SQLiteRepository) GetAll() ([]domain.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at, tags FROM notes ORDER BY updated_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []domain.Note
	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func (r *SQLiteRepository) GetByID(id int) (*domain.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at, tags FROM notes WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var note domain.Note
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *SQLiteRepository) Update(note *domain.Note) error {
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

func (r *SQLiteRepository) Search(query string) ([]domain.Note, error) {
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

	var notes []domain.Note
	for rows.Next() {
		var note domain.Note
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

// UpdateEmbedding stores the embedding vector for a note
func (r *SQLiteRepository) UpdateEmbedding(noteID int, embedding string) error {
	query := `UPDATE notes SET embedding = ? WHERE id = ?`
	_, err := r.db.Exec(query, embedding, noteID)
	return err
}

// GetNotesWithEmbeddings retrieves all notes that have embeddings
func (r *SQLiteRepository) GetNotesWithEmbeddings() ([]domain.Note, error) {
	query := `SELECT id, title, content, created_at, updated_at, tags, COALESCE(embedding, '') as embedding
			  FROM notes WHERE embedding IS NOT NULL ORDER BY updated_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []domain.Note
	for rows.Next() {
		var note domain.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Tags, &note.Embedding)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

// SaveChatMessage stores a chat message
func (r *SQLiteRepository) SaveChatMessage(role, content string) error {
	query := `INSERT INTO chat_history (role, content, created_at) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, role, content, time.Now())
	return err
}

// GetChatHistory retrieves recent chat messages
func (r *SQLiteRepository) GetChatHistory(limit int) ([]domain.ChatMessage, error) {
	query := `SELECT id, role, content, created_at FROM chat_history ORDER BY created_at DESC LIMIT ?`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []domain.ChatMessage
	for rows.Next() {
		var msg domain.ChatMessage
		err := rows.Scan(&msg.ID, &msg.Role, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append([]domain.ChatMessage{msg}, messages...) // Reverse order
	}

	return messages, rows.Err()
}

// ClearChatHistory removes all chat messages
func (r *SQLiteRepository) ClearChatHistory() error {
	query := `DELETE FROM chat_history`
	_, err := r.db.Exec(query)
	return err
}
