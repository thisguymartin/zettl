package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"thisguymartin/zettl/internal/infrastructure/database"
	ui "thisguymartin/zettl/internal/ui"
)

func main() {
	repo, err := database.NewSQLiteRepository("internal/infrastructure/database/zettl.db")
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	model, err := ui.NewUIModel(repo)
	if err != nil {
		log.Fatalf("failed to initialize UI model: %v", err)
	}

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
