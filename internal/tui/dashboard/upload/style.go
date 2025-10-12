package upload

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var Info = style.Info.Padding(0, 25)

var Directory = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Primary)).
	Bold(true)

var Selected = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Secondary)).
	Bold(true)

var Cursor = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Secondary)).
	Bold(true)

var File = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Foregrounds)).
	Bold(true)

var Succes = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Tertiary)).
	Bold(true)
