package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	//zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
	"os"
	internal "thisguymartin/zettl/internal"
	"thisguymartin/zettl/internal/infrastructure/database"
)

var (
	// cfgFile string

	rootCmd = &cobra.Command{
		Use:     "gh dash",
		Short:   "A gh extension that shows a configurable dashboard of pull requests and issues.",
		Version: "",
		Args:    cobra.MaximumNArgs(1),
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	repo, err := database.NewSQLiteRepository("internal/infrastructure/database/zettl.db")
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	//zone.NewGlobal()

	model, _ := internal.New(repo)
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	Execute()
}
