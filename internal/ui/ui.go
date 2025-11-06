package domain

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type WindowType int

const (
	MainMenuWindow WindowType = iota
	NoteListWindow
	NoteEditWindow
	SearchWindow
	ChatWindow
	AIMenuWindow
)

type UIModel struct {
	currentWindow WindowType
	notes         []Note
	filteredNotes []Note
	cursor        int
	searchQuery   string
	noteContent   string
	noteTitle     string
	selectedNote  *Note
	repo          NoteRepository
	width         int
	height        int
	// AI-related fields
	chatInput    string
	chatHistory  []ChatMessage
	isProcessing bool
	errorMsg     string
	aiService    interface{} // Will hold *ai.AIService
	// Bubbles components
	textarea    textarea.Model
	textinput   textinput.Model
	chatinput   textinput.Model
	searchinput textinput.Model
	spinner     spinner.Model
	notesList   list.Model
}

func NewUIModel(repo NoteRepository) (*UIModel, error) {
	notes, err := repo.GetAll()
	if err != nil {
		log.Error("failed to get notes: %v", "error", err)
		notes = []Note{}
	}

	// Initialize textarea for note editing
	ta := textarea.New()
	ta.Placeholder = "Start typing your note..."
	ta.Focus()
	ta.CharLimit = 10000
	ta.SetWidth(80)
	ta.SetHeight(20)
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(true)

	// Initialize text input for title
	ti := textinput.New()
	ti.Placeholder = "Note title (optional)"
	ti.CharLimit = 200
	ti.Width = 80

	// Initialize chat input
	ci := textinput.New()
	ci.Placeholder = "Type your message..."
	ci.CharLimit = 500
	ci.Width = 80

	// Initialize search input
	si := textinput.New()
	si.Placeholder = "Search notes..."
	si.CharLimit = 100
	si.Width = 80
	si.Focus()

	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &UIModel{
		currentWindow: MainMenuWindow,
		notes:         notes,
		filteredNotes: notes,
		repo:          repo,
		width:         80,
		height:        24,
		chatHistory:   []ChatMessage{},
		textarea:      ta,
		textinput:     ti,
		chatinput:     ci,
		searchinput:   si,
		spinner:       s,
	}, nil
}

// SetAIService sets the AI service for the model
func (m *UIModel) SetAIService(service interface{}) {
	m.aiService = service
}

func (m UIModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		textarea.Blink,
	)
}

func (m UIModel) View() string {
	switch m.currentWindow {
	case MainMenuWindow:
		return m.viewMainMenu()
	case NoteListWindow:
		return m.viewNoteList()
	case NoteEditWindow:
		return m.viewNoteEdit()
	case SearchWindow:
		return m.viewSearch()
	case ChatWindow:
		return m.viewChat()
	case AIMenuWindow:
		return m.viewAIMenu()
	}
	return ""
}

func (m UIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.textarea.SetWidth(msg.Width - 8)
		m.textarea.SetHeight(msg.Height - 12)
		return m, nil

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		// Update bubbles components based on current window
		switch m.currentWindow {
		case NoteEditWindow:
			m.textarea, cmd = m.textarea.Update(msg)
			cmds = append(cmds, cmd)
		case SearchWindow:
			m.searchinput, cmd = m.searchinput.Update(msg)
			cmds = append(cmds, cmd)
		case ChatWindow:
			m.chatinput, cmd = m.chatinput.Update(msg)
			cmds = append(cmds, cmd)
		}

		// Handle window-specific key bindings
		switch m.currentWindow {
		case MainMenuWindow:
			return m.updateMainMenu(msg)
		case NoteListWindow:
			return m.updateNoteList(msg)
		case NoteEditWindow:
			return m.updateNoteEdit(msg)
		case SearchWindow:
			return m.updateSearch(msg)
		case ChatWindow:
			return m.updateChat(msg)
		case AIMenuWindow:
			return m.updateAIMenu(msg)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *UIModel) refreshNotes() {
	notes, err := m.repo.GetAll()
	if err != nil {
		log.Printf("failed to refresh notes: %v", err)
		return
	}
	m.notes = notes
	m.applyFilter()
}

func (m *UIModel) applyFilter() {
	if m.searchQuery == "" {
		m.filteredNotes = m.notes
		return
	}

	filtered := []Note{}
	query := strings.ToLower(m.searchQuery)
	for _, note := range m.notes {
		if strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Content), query) ||
			strings.Contains(strings.ToLower(note.Tags), query) {
			filtered = append(filtered, note)
		}
	}
	m.filteredNotes = filtered
}

func (m UIModel) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "1", "n":
		m.currentWindow = NoteEditWindow
		m.noteTitle = ""
		m.noteContent = ""
		m.selectedNote = nil
	case "2", "l":
		m.currentWindow = NoteListWindow
		m.cursor = 0
	case "3", "s":
		m.currentWindow = SearchWindow
		m.searchQuery = ""
		m.cursor = 0
	case "4", "a":
		m.currentWindow = AIMenuWindow
		m.cursor = 0
	}
	return m, nil
}

func (m UIModel) updateNoteList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		m.currentWindow = MainMenuWindow
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.filteredNotes)-1 {
			m.cursor++
		}
	case "enter":
		if len(m.filteredNotes) > 0 && m.cursor < len(m.filteredNotes) {
			m.selectedNote = &m.filteredNotes[m.cursor]
			m.noteTitle = m.selectedNote.Title
			m.noteContent = m.selectedNote.Content
			m.currentWindow = NoteEditWindow
		}
	case "/":
		m.currentWindow = SearchWindow
		m.searchQuery = ""
	}
	return m, nil
}

func (m UIModel) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		m.currentWindow = NoteListWindow
		m.searchQuery = ""
		m.searchinput.SetValue("")
		m.applyFilter()
	case "enter":
		m.searchQuery = m.searchinput.Value()
		m.currentWindow = NoteListWindow
		m.cursor = 0
		m.applyFilter()
	}
	// Sync search query with textinput value and apply filter on every change
	m.searchQuery = m.searchinput.Value()
	m.applyFilter()
	return m, nil
}

func (m UIModel) updateNoteEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = NoteListWindow
		m.noteContent = "" // Clear on cancel
		m.textarea.Reset()
	case "ctrl+s":
		// Get content from textarea before saving
		m.noteContent = m.textarea.Value()
		return m.saveNote()
	}
	// Textarea handles all other input via Update() in the main Update function
	// Just sync the noteContent with textarea value
	m.noteContent = m.textarea.Value()
	return m, nil
}

func (m UIModel) saveNote() (tea.Model, tea.Cmd) {
	if m.noteContent == "" {
		return m, nil
	}

	if m.noteTitle == "" {
		lines := strings.Split(m.noteContent, "\n")
		if len(lines) > 0 && strings.TrimSpace(lines[0]) != "" {
			m.noteTitle = strings.TrimSpace(lines[0])
			if len(m.noteTitle) > 50 {
				m.noteTitle = m.noteTitle[:50] + "..."
			}
		} else {
			m.noteTitle = "Untitled Note"
		}
	}

	if m.selectedNote != nil {
		m.selectedNote.Title = m.noteTitle
		m.selectedNote.Content = m.noteContent
		m.selectedNote.UpdatedAt = time.Now()
		err := m.repo.Update(m.selectedNote)
		if err != nil {
			log.Printf("failed to update note: %v", err)
		}
	} else {
		note := &Note{
			Title:     m.noteTitle,
			Content:   m.noteContent,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Tags:      "notebook",
		}
		err := m.repo.Create(note)
		if err != nil {
			log.Printf("failed to create note: %v", err)
		}
	}

	m.refreshNotes()
	m.currentWindow = NoteListWindow
	return m, nil
}

func (m UIModel) viewMainMenu() string {
	// Modern logo with gradient effect
	logo := logoStyle.Render(`
╭─────────────────────────────────────────────╮
│                                             │
│     ███████╗███████╗████████╗████████╗██╗  │
│     ╚══███╔╝██╔════╝╚══██╔══╝╚══██╔══╝██║  │
│       ███╔╝ █████╗     ██║      ██║   ██║  │
│      ███╔╝  ██╔══╝     ██║      ██║   ██║  │
│     ███████╗███████╗   ██║      ██║   ███╗ │
│     ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══╝ │
│                                             │
╰─────────────────────────────────────────────╯`)

	title := bigTitleStyle.Width(m.width - 4).Align(lipgloss.Center).Render("✨ AI-Powered Note Taking ✨")

	// Menu items with modern styling
	menuItems := []string{
		renderKeyHelp("n/1", "New Note") + "     " + infoStyle.Render("Create a new note"),
		renderKeyHelp("l/2", "List Notes") + "   " + infoStyle.Render("Browse existing notes"),
		renderKeyHelp("s/3", "Search") + "       " + infoStyle.Render("Find notes instantly"),
		renderKeyHelp("a/4", "AI Features") + "  " + successStyle.Render("✨ Chat & Semantic Search"),
	}

	menuCard := cardStyle.
		Width(m.width - 6).
		Align(lipgloss.Center).
		Render(strings.Join(menuItems, "\n\n"))

	// Help text
	help := helpStyle.Render("Press the corresponding key to navigate • Press 'q' to quit")

	return lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(lipgloss.JoinVertical(lipgloss.Center,
			"",
			logo,
			"",
			title,
			"",
			menuCard,
			"",
			help,
		))
}

func (m UIModel) viewNoteList() string {
	header := headerStyle.Render(fmt.Sprintf("📚 Notes %s", renderBadge(fmt.Sprintf("%d", len(m.filteredNotes)))))

	if len(m.filteredNotes) == 0 {
		empty := cardStyle.
			Width(m.width - 6).
			Align(lipgloss.Center).
			Render(warningStyle.Render("No notes found") + "\n\n" + helpStyle.Render("Press 'Esc' to go back or '/' to search"))
		return lipgloss.JoinVertical(lipgloss.Left, header, "", empty)
	}

	var notes []string
	visibleStart := 0
	visibleEnd := len(m.filteredNotes)
	maxVisible := (m.height - 10) / 6 // Approximate notes that fit

	if maxVisible > 0 && len(m.filteredNotes) > maxVisible {
		if m.cursor > maxVisible/2 {
			visibleStart = m.cursor - maxVisible/2
		}
		visibleEnd = visibleStart + maxVisible
		if visibleEnd > len(m.filteredNotes) {
			visibleEnd = len(m.filteredNotes)
			visibleStart = visibleEnd - maxVisible
			if visibleStart < 0 {
				visibleStart = 0
			}
		}
	}

	for i := visibleStart; i < visibleEnd; i++ {
		note := m.filteredNotes[i]
		dateStr := note.CreatedAt.Format("Jan 2, 2006 15:04")
		preview := strings.ReplaceAll(note.Content, "\n", " ")
		if len(preview) > 80 {
			preview = preview[:80] + "..."
		}

		titleLine := titleStyle.Render("📄 " + note.Title)
		dateLine := helpStyle.Render("🕒 " + dateStr)
		previewLine := infoStyle.Render("💭 " + preview)

		content := lipgloss.JoinVertical(lipgloss.Left, titleLine, dateLine, previewLine)

		if i == m.cursor {
			notes = append(notes, activeCardStyle.Width(m.width-6).Render(content))
		} else {
			notes = append(notes, cardStyle.Width(m.width-6).Render(content))
		}
	}

	// Help footer
	helpKeys := []string{
		renderKeyHelp("↑↓/jk", "Navigate"),
		renderKeyHelp("Enter", "Open"),
		renderKeyHelp("/", "Search"),
		renderKeyHelp("Esc", "Back"),
	}
	footer := helpStyle.Render(strings.Join(helpKeys, " • "))

	return lipgloss.JoinVertical(lipgloss.Left, header, "", strings.Join(notes, "\n"), "", footer)
}

func (m UIModel) viewSearch() string {
	header := headerStyle.Render("🔍 Search Notes")

	// Update search input value
	m.searchinput.SetValue(m.searchQuery)
	m.searchinput.Width = m.width - 10

	searchBox := focusedInputStyle.
		Width(m.width - 6).
		Render(m.searchinput.View())

	resultsText := ""
	if len(m.searchQuery) > 0 {
		if len(m.filteredNotes) > 0 {
			resultsText = successStyle.Render(fmt.Sprintf("✓ Found %d notes", len(m.filteredNotes)))
		} else {
			resultsText = warningStyle.Render("No notes found")
		}
	} else {
		resultsText = helpStyle.Render("Start typing to search...")
	}

	// Help footer
	helpKeys := []string{
		renderKeyHelp("Type", "Search"),
		renderKeyHelp("Enter", "View Results"),
		renderKeyHelp("Esc", "Back"),
	}
	footer := helpStyle.Render(strings.Join(helpKeys, " • "))

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		searchBox,
		"",
		resultsText,
		"",
		footer,
	)
}

func (m UIModel) viewNoteEdit() string {
	var titleText string
	if m.selectedNote != nil {
		titleText = fmt.Sprintf("✏️  Editing: %s", renderBadge(m.selectedNote.Title))
	} else {
		titleText = "✏️  " + successStyle.Render("New Note")
	}

	header := headerStyle.Render(titleText)

	// Update textarea with current content
	if m.textarea.Value() != m.noteContent {
		m.textarea.SetValue(m.noteContent)
	}

	// Style the textarea container
	editorBox := focusedInputStyle.
		Width(m.width - 6).
		Render(m.textarea.View())

	// Help footer with modern key hints
	helpKeys := []string{
		renderKeyHelp("Ctrl+S", "Save"),
		renderKeyHelp("Esc", "Cancel"),
		renderKeyHelp("Tab", "Indent"),
	}
	footer := helpStyle.Render(strings.Join(helpKeys, " • "))

	// Word count
	wordCount := len(strings.Fields(m.noteContent))
	charCount := len(m.noteContent)
	stats := helpStyle.Render(fmt.Sprintf("📊 %d words • %d characters", wordCount, charCount))

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		editorBox,
		"",
		stats,
		footer,
	)
}

func (m UIModel) viewAIMenu() string {
	header := headerStyle.Render("🤖 AI Features " + successStyle.Render("✨"))

	// Menu items with modern styling
	menuItems := []string{
		renderKeyHelp("c/1", "Chat") + "              " + infoStyle.Render("Chat with AI about your notes"),
		renderKeyHelp("2", "Generate Embeddings") + " " + infoStyle.Render("Create vectors for semantic search"),
		renderKeyHelp("3", "Semantic Search") + "    " + infoStyle.Render("Find similar notes using AI"),
		renderKeyHelp("4", "Clear Chat History") + " " + warningStyle.Render("Reset conversation"),
	}

	menuCard := cardStyle.
		Width(m.width - 6).
		Align(lipgloss.Left).
		Render(strings.Join(menuItems, "\n\n"))

	// Status message if any
	statusMsg := ""
	if m.errorMsg != "" {
		if strings.Contains(m.errorMsg, "cleared") || strings.Contains(m.errorMsg, "success") {
			statusMsg = successStyle.Render("✓ " + m.errorMsg)
		} else if strings.Contains(m.errorMsg, "coming soon") {
			statusMsg = infoStyle.Render("ℹ " + m.errorMsg)
		} else {
			statusMsg = errorStyle.Render("✗ " + m.errorMsg)
		}
	}

	// Help footer
	footer := helpStyle.Render("Press the number or letter key • Press 'Esc' to go back")

	parts := []string{header, "", menuCard}
	if statusMsg != "" {
		parts = append(parts, "", statusMsg)
	}
	parts = append(parts, "", footer)

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

func (m UIModel) updateAIMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = MainMenuWindow
		m.errorMsg = ""
	case "1", "c":
		// Load chat history
		if history, err := m.repo.GetChatHistory(50); err == nil {
			m.chatHistory = history
		}
		m.currentWindow = ChatWindow
		m.chatInput = ""
		m.errorMsg = ""
	case "2":
		// Generate embeddings - to be implemented
		m.errorMsg = "Embedding generation - coming soon!"
	case "3":
		// Semantic search - to be implemented
		m.errorMsg = "Semantic search - coming soon!"
	case "4":
		// Clear chat history
		if err := m.repo.ClearChatHistory(); err == nil {
			m.chatHistory = []ChatMessage{}
			m.errorMsg = "Chat history cleared!"
		} else {
			m.errorMsg = "Failed to clear chat history"
		}
	}
	return m, nil
}

func (m UIModel) viewChat() string {
	header := headerStyle.Render("💬 AI Chat " + renderBadge(fmt.Sprintf("%d msgs", len(m.chatHistory))))

	// Render chat history with modern styling
	var chatMessages []string
	if len(m.chatHistory) == 0 {
		emptyMsg := cardStyle.
			Width(m.width - 6).
			Align(lipgloss.Center).
			Render(infoStyle.Render("No messages yet. Start a conversation!") + "\n\n" +
				helpStyle.Render("💡 Tip: The AI has access to your notes for context"))
		chatMessages = append(chatMessages, emptyMsg)
	} else {
		// Show last N messages that fit
		maxMessages := (m.height - 15) / 4
		startIdx := 0
		if len(m.chatHistory) > maxMessages {
			startIdx = len(m.chatHistory) - maxMessages
		}

		for i := startIdx; i < len(m.chatHistory); i++ {
			msg := m.chatHistory[i]
			var msgStyle lipgloss.Style
			var roleText string

			if msg.Role == "assistant" {
				msgStyle = cardStyle.BorderForeground(colorSecondary)
				roleText = successStyle.Render("🤖 AI")
			} else {
				msgStyle = cardStyle.BorderForeground(colorPrimary)
				roleText = titleStyle.Render("👤 You")
			}

			content := lipgloss.JoinVertical(lipgloss.Left,
				roleText,
				helpStyle.Render(msg.Content),
			)

			chatMessages = append(chatMessages, msgStyle.Width(m.width-6).Render(content))
		}
	}

	// Processing indicator
	if m.isProcessing {
		processingMsg := cardStyle.
			Width(m.width - 6).
			Render(m.spinner.View() + " " + infoStyle.Render("AI is thinking..."))
		chatMessages = append(chatMessages, processingMsg)
	}

	// Update chat input
	m.chatinput.SetValue(m.chatInput)
	m.chatinput.Width = m.width - 10

	inputBox := focusedInputStyle.
		Width(m.width - 6).
		Render(m.chatinput.View())

	// Help footer
	helpKeys := []string{
		renderKeyHelp("Type", "Message"),
		renderKeyHelp("Enter", "Send"),
		renderKeyHelp("Esc", "Back"),
	}
	footer := helpStyle.Render(strings.Join(helpKeys, " • "))

	// Status message
	statusMsg := ""
	if m.errorMsg != "" {
		statusMsg = errorStyle.Render("✗ " + m.errorMsg)
	}

	parts := []string{header, "", strings.Join(chatMessages, "\n"), "", inputBox}
	if statusMsg != "" {
		parts = append(parts, "", statusMsg)
	}
	parts = append(parts, footer)

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

func (m UIModel) updateChat(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.currentWindow = AIMenuWindow
		m.chatInput = ""
		m.errorMsg = ""
	case "enter":
		m.chatInput = m.chatinput.Value()
		if m.chatInput != "" && !m.isProcessing {
			// Save user message
			userMsg := m.chatInput
			if err := m.repo.SaveChatMessage("user", userMsg); err != nil {
				m.errorMsg = "Failed to save message"
				return m, nil
			}

			// Add to history
			m.chatHistory = append(m.chatHistory, ChatMessage{
				Role:    "user",
				Content: userMsg,
			})

			m.chatInput = ""
			m.chatinput.SetValue("")
			m.isProcessing = true
			m.errorMsg = ""

			// TODO: In a real implementation, this would call the AI service asynchronously
			// For now, just add a placeholder response
			response := "AI integration requires OpenAI API key. Please set OPENAI_API_KEY environment variable or configure it in ~/.zettl/config.json"

			if err := m.repo.SaveChatMessage("assistant", response); err != nil {
				m.errorMsg = "Failed to save AI response"
			} else {
				m.chatHistory = append(m.chatHistory, ChatMessage{
					Role:    "assistant",
					Content: response,
				})
			}
			m.isProcessing = false
		}
	}
	// Sync chatInput with textinput value
	m.chatInput = m.chatinput.Value()
	return m, nil
}
