package main

import (
	"os"
	"thisguymartin/zettl/internal/infrastructure/database"
	ui "thisguymartin/zettl/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
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

func createModel(db *database.SQLiteRepository, debug bool) (*ui.UIModel, error) {
	// if debug {
	// 	var fileErr error
	// 	newConfigFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// 	if fileErr == nil {
	// 		log.SetOutput(newConfigFile)
	// 		log.SetTimeFormat(time.Kitchen)
	// 		log.SetReportCaller(true)
	// 		log.SetLevel(log.DebugLevel)
	// 		log.Debug("Logging to debug.log")
	// 		if repoPath != "" {
	// 			log.Debug("Running in repo", "repo", repoPath)
	// 		}
	// 	} else {
	// 		loggerFile, _ = tea.LogToFile("debug.log", "debug")
	// 		slog.Print("Failed setting up logging", fileErr)
	// 	}
	// } else {
	// 	log.SetOutput(os.Stderr)
	// 	log.SetLevel(log.FatalLevel)
	// }

	return ui.NewUIModel(db)
}

func init() {

	repo, err := database.NewSQLiteRepository("internal/infrastructure/database/zettl.db")
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	zone.NewGlobal()

	model, _ := createModel(repo, false)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	Execute()
}
