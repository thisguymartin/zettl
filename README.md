# Zettl

Zettl is a terminal-based note-taking CLI application built with Go, using the Bubble Tea TUI framework and a SQLite database for persistent storage.

## Features
- View your notes in a terminal UI
- Add new notes interactively (press `n` to start typing)
- Notes are saved to a local SQLite database (`zettl.db`)
- Navigate notes with arrow keys or `j`/`k`
- Select/deselect notes with space or enter
- Quit with `q` or `ctrl+c`

## Usage

### Run the App
```sh
go run ./cmd/main.go
```

### Controls
- `n` — New note (type your note, press Enter to save, Esc to cancel)
- `up`/`k` — Move cursor up
- `down`/`j` — Move cursor down
- `space`/`enter` — Select/deselect a note
- `q` or `ctrl+c` — Quit

## Database
- Notes are stored in `zettl.db` in the project root.
- Each note has a title, content, tags, and timestamps.

## Dependencies
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)

Install dependencies with:
```sh
go mod tidy
```

## Project Structure
- `cmd/main.go` — Main TUI application
- `internal/domain/note.go` — Note model and repository interface
- `internal/domain/database.go` — SQLite repository implementation

## License
MIT
