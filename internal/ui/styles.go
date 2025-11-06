package domain

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Simple color palette
	colorPrimary = lipgloss.Color("205") // Purple
	colorMuted   = lipgloss.Color("240") // Gray
	colorAccent  = lipgloss.Color("212") // Pink

	// Base styles
	titleStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	selectedStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Padding(1)
)
