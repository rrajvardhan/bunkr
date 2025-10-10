package auth

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Padding(0, 9).
			Foreground(lipgloss.Color(theme.Colors.Primary))

	DividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.SelectionsNoAlpha))

	ActiveChoice = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.Secondary)).
			Bold(true)

	InactiveChoice = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.Comments))

	Info = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Comments)).
		Padding(0, 9).
		Faint(true)

	Subtle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Quinary)).
		Faint(true)

	Cursor = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Quinary))

	Error = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Senary)).
		Bold(true)
)
