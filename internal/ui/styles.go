package domain

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	colorPrimary   = lipgloss.Color("#7D56F4")
	colorSecondary = lipgloss.Color("#9D7FF5")
	colorAccent    = lipgloss.Color("#FF6B9D")
	colorSuccess   = lipgloss.Color("#32D74B")
	colorWarning   = lipgloss.Color("#FFD60A")
	colorDanger    = lipgloss.Color("#FF453A")
	colorBg        = lipgloss.Color("#1a1a1a")
	colorBgAlt     = lipgloss.Color("#2a2a2a")
	colorText      = lipgloss.Color("#FAFAFA")
	colorTextMuted = lipgloss.Color("#A0A0A0")
	colorBorder    = lipgloss.Color("#404040")

	// Base styles
	baseStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Background(colorBg)

	// Title styles with gradient effect
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Padding(0, 1).
			MarginBottom(1)

	bigTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Padding(1, 2).
			MarginBottom(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			BorderBottom(true)

	// Header styles
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorText).
			Background(colorPrimary).
			Padding(0, 2).
			MarginBottom(1)

	// Card styles with shadow effect
	cardStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1, 2).
			MarginBottom(1)

	activeCardStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Background(colorBgAlt).
			Padding(1, 2).
			MarginBottom(1).
			Bold(true)

	// Input styles
	inputStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(colorSecondary).
			Padding(1, 2).
			MarginBottom(1)

	focusedInputStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(colorPrimary).
				Padding(1, 2).
				MarginBottom(1)

	// Status styles
	successStyle = lipgloss.NewStyle().
			Foreground(colorSuccess).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorDanger).
			Bold(true)

	warningStyle = lipgloss.NewStyle().
			Foreground(colorWarning).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(colorSecondary).
			Bold(true)

	// Menu item styles
	menuItemStyle = lipgloss.NewStyle().
			Foreground(colorTextMuted).
			Padding(0, 2)

	selectedMenuItemStyle = lipgloss.NewStyle().
				Foreground(colorPrimary).
				Bold(true).
				Padding(0, 2).
				Background(colorBgAlt)

	// Footer/help text style
	helpStyle = lipgloss.NewStyle().
			Foreground(colorTextMuted).
			MarginTop(1)

	// Divider
	dividerStyle = lipgloss.NewStyle().
			Foreground(colorBorder).
			MarginTop(1).
			MarginBottom(1)

	// Badge styles
	badgeStyle = lipgloss.NewStyle().
			Background(colorPrimary).
			Foreground(colorText).
			Padding(0, 1).
			MarginRight(1).
			Bold(true)

	// Logo style with gradient
	logoStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true)

	// Content area
	contentStyle = lipgloss.NewStyle().
			Padding(1, 2).
			MarginTop(1)
)

// Helper functions for dynamic styling
func renderDivider(width int) string {
	return dividerStyle.Render(strings.Repeat("─", width-4))
}

func renderBadge(text string) string {
	return badgeStyle.Render(text)
}

func renderKeyHelp(key, desc string) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(colorPrimary).
		Bold(true).
		Padding(0, 1).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(colorPrimary)

	descStyle := lipgloss.NewStyle().
		Foreground(colorTextMuted)

	return lipgloss.JoinHorizontal(lipgloss.Center,
		keyStyle.Render(key),
		descStyle.Render(" "+desc),
	)
}
