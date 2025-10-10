package dashboard

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
		Faint(true)

	Subtle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Quinary)).
		Faint(true)
)
