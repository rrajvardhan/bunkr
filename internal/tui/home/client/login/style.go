package login

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var Error = lipgloss.NewStyle().
	Foreground(lipgloss.Color(theme.Colors.Senary)).
	Bold(true)

var (
	Info   = style.Info.Padding(0)
	Header = style.Header.Padding(0)
)
