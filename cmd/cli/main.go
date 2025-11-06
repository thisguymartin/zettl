package main

import (
	"fmt"
	"os"
	"thisguymartin/zettl/internal/config"
	"thisguymartin/zettl/internal/infrastructure/ai"
	"thisguymartin/zettl/internal/infrastructure/database"
	ui "thisguymartin/zettl/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "zettl",
		Short:   "A powerful note-taking app with AI features",
		Version: "1.0.0",
		Long: `Zettl is a terminal-based note-taking application with AI-powered features including:
- Full CRUD operations for notes
- Semantic search using embeddings
- AI chat with context from your notes
- Export/import functionality`,
		Args: cobra.MaximumNArgs(1),
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func createModel(repo *database.SQLiteRepository, aiService *ai.AIService) (*ui.UIModel, error) {
	model, err := ui.NewUIModel(repo)
	if err != nil {
		return nil, err
	}

	// Inject AI service if available
	if aiService != nil {
		// Store as interface to avoid circular dependency
		model.SetAIService(aiService)
	}

	return model, nil
}

func init() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Warnf("failed to load config: %v (using defaults)", err)
		cfg = config.DefaultConfig()
	}

	// Print config location for user reference
	if cfg.OpenAIAPIKey == "" {
		fmt.Printf("💡 Tip: Set OPENAI_API_KEY or configure it in %s to enable AI features\n\n", config.ConfigPath())
	}

	// Initialize database with configured path
	repo, err := database.NewSQLiteRepository(cfg.DBPath)
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	// Initialize AI service if API key is available
	var aiService *ai.AIService
	if cfg.OpenAIAPIKey != "" {
		aiService = ai.NewAIService(cfg.OpenAIAPIKey)
		log.Info("AI features enabled")
	} else {
		log.Info("AI features disabled (no API key)")
	}

	zone.NewGlobal()

	model, err := createModel(repo, aiService)
	if err != nil {
		log.Fatalf("failed to create model: %v", err)
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Execute()
}
