package files

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var Info = style.Info.Padding(0, 17)

var Succes = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Tertiary)).
	Bold(true)
