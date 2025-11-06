# Zettl - AI-Powered Note-Taking App

A powerful terminal-based note-taking application with AI features including semantic search, embeddings, and chat capabilities.

## Features

### Core Features
- ✅ **Full CRUD Operations** - Create, read, update, and delete notes
- ✅ **SQLite Persistence** - All notes stored locally in SQLite database
- ✅ **Bubble Tea TUI** - Beautiful terminal user interface
- ✅ **Tag System** - Organize notes with tags
- ✅ **Fuzzy Search** - Find notes by title, content, or tags

### AI-Powered Features (NEW!)
- 🤖 **AI Chat** - Conversational AI with context from your notes
- 🔍 **Semantic Search** - Find similar notes using vector embeddings
- 📊 **Embedding Generation** - Create vector representations of notes
- 💾 **Chat History** - Persistent conversation storage
- 🎯 **Context-Aware** - AI has access to your notes for better responses

### Import/Export
- 📤 **JSON Export** - Export all notes to structured JSON
- 📥 **JSON Import** - Import notes from JSON backup
- 📝 **Markdown Export** - Export notes as individual markdown files with frontmatter

## Installation

### Prerequisites
- Go 1.24.3 or later
- OpenAI API key (optional, for AI features)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/thisguymartin/zettl.git
cd zettl

# Build the application
make build

# Or run directly
make dev
```

## Configuration

### Setting up OpenAI API Key

Zettl supports two methods for configuring your OpenAI API key:

#### Method 1: Environment Variable (Recommended)
```bash
export OPENAI_API_KEY="sk-your-api-key-here"
```

#### Method 2: Configuration File
1. Run the app once to create the config directory
2. Edit `~/.zettl/config.json`:
```json
{
  "openai_api_key": "sk-your-api-key-here",
  "db_path": "/home/user/.zettl/zettl.db"
}
```

### Database Location
By default, the database is stored at `~/.zettl/zettl.db`. You can change this in the config file.

## Usage

### Starting the App
```bash
# Run the compiled binary
./bin/cli

# Or using make
make dev

# Or with hot reload (development)
make run
```

### Navigation

#### Main Menu
- `1` or `n` - Create a new note
- `2` or `l` - List all notes
- `3` or `s` - Search notes
- `4` or `a` - AI Features menu
- `q` or `Ctrl+C` - Quit

#### Note List
- `↑/↓` or `j/k` - Navigate through notes
- `Enter` - Open selected note
- `/` - Quick search
- `Esc` - Return to main menu

#### Note Editor
- Type to add content
- `Enter` - New line
- `Tab` - Indent (4 spaces)
- `Ctrl+S` - Save note
- `Esc` - Cancel and return

#### AI Features Menu
- `1` or `c` - Open chat interface
- `2` - Generate embeddings for all notes
- `3` - Semantic search using embeddings
- `4` - Clear chat history
- `Esc` - Return to main menu

#### Chat Window
- Type your message
- `Enter` - Send message
- `Esc` - Return to AI menu

## Architecture

### Domain Model
```
zettl/
├── cmd/cli/                    # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   │   └── config.go
│   ├── infrastructure/
│   │   ├── ai/                 # AI service layer
│   │   │   └── ai_service.go   # OpenAI integration
│   │   ├── database/           # Data persistence
│   │   │   └── sqlite.go       # SQLite repository
│   │   └── export/             # Import/export functionality
│   │       └── export.go
│   └── ui/                     # User interface
│       ├── note.go             # Domain models
│       └── ui.go               # Bubble Tea TUI
├── go.mod
└── Makefile
```

### Technology Stack

**Core:**
- **Go 1.24.3** - Programming language
- **SQLite** - Local database (via go-libsql)
- **Bubble Tea** - Terminal UI framework
- **Lip Gloss** - Terminal styling

**AI Integration:**
- **OpenAI API** - Embeddings and chat completions
- **text-embedding-3-small** - Embedding model (1536 dimensions)
- **gpt-4o-mini** - Chat model

## AI Features Deep Dive

### How Embeddings Work

1. **Generation**: When you generate embeddings, each note's title and content are converted into a 1536-dimensional vector
2. **Storage**: Vectors are stored as JSON in the SQLite database
3. **Search**: Query text is converted to a vector and compared using cosine similarity
4. **Results**: Most similar notes are returned, even if they don't share keywords

### Chat with Context

1. **Context Retrieval**: When you chat, the AI can access your notes
2. **Semantic Search**: Relevant notes are found based on your question
3. **Enhanced Responses**: AI uses note content to provide informed answers
4. **Persistent History**: Conversations are saved across sessions

### Cost Estimation

**Embeddings:**
- ~$0.0001 per 1,000 tokens
- Average note: ~100-500 tokens
- 1,000 notes: ~$0.01-0.05 one-time cost

**Chat:**
- ~$0.01-0.03 per conversation
- Context-aware responses use more tokens
- Average conversation (10 messages): ~$0.05-0.15

**Monthly Costs (Example):**
- 1,000 notes with embeddings: $0.05 (one-time)
- 100 chat conversations/month: $5-15
- Total: ~$5-15/month

### Local Alternative
For a free, local alternative:
- Replace OpenAI with local embedding models (all-MiniLM-L6-v2)
- Use local LLMs (Ollama, LM Studio)
- Trade-off: Slower performance, requires GPU for best results

## Database Schema

### Notes Table
```sql
CREATE TABLE notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    tags TEXT DEFAULT '',
    embedding TEXT DEFAULT NULL  -- JSON-encoded vector
);
```

### Chat History Table
```sql
CREATE TABLE chat_history (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role TEXT NOT NULL,          -- 'user' or 'assistant'
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Development

### Project Structure
```
internal/
├── config/              # Configuration management
├── infrastructure/
│   ├── ai/             # AI service integration
│   ├── database/       # Data persistence layer
│   └── export/         # Import/export utilities
└── ui/                 # User interface layer
    ├── note.go         # Domain models and interfaces
    └── ui.go           # Bubble Tea UI implementation
```

### Adding New Features

1. **New AI Models**: Update `ai_service.go` constants
2. **New Windows**: Add to `WindowType` enum in `ui.go`
3. **New Repository Methods**: Update interface in `note.go` and implementation in `sqlite.go`

### Testing

```bash
# Run tests (when available)
go test ./...

# Build for production
make build

# Run with hot reload
make run
```

## Troubleshooting

### AI Features Not Working
- Check if `OPENAI_API_KEY` is set correctly
- Verify API key has credits
- Check network connectivity
- Review error messages in the UI

### Database Errors
- Check file permissions for `~/.zettl/`
- Ensure sufficient disk space
- Verify SQLite installation

### Build Issues
- Ensure Go 1.24.3+ is installed
- Run `go mod tidy` to sync dependencies
- Check for missing system libraries

## Roadmap

### Phase 1 ✅ (Completed)
- Basic CRUD operations
- SQLite integration
- Bubble Tea TUI

### Phase 2 ✅ (Completed)
- Tag system
- Search functionality
- Export/import

### Phase 3 🚀 (Current)
- ✅ Embedding generation
- ✅ Vector similarity search
- ✅ Chat interface
- ⏳ Async AI operations

### Phase 4 📋 (Planned)
- Advanced chat features
- Local LLM support
- Performance optimization
- Cross-platform packaging

## Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

[Your chosen license]

## Acknowledgments

- [Charm Bracelet](https://charm.sh/) - Bubble Tea and Lip Gloss
- [OpenAI](https://openai.com/) - AI capabilities
- [Turso](https://turso.tech/) - SQLite driver

## Support

For issues, questions, or suggestions:
- Open an issue on GitHub
- Check existing documentation
- Review troubleshooting section

---

Built with ❤️ using Go and Bubble Tea
