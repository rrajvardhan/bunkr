package dashboard

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var Info = style.Info.
	Padding(0, 2)

var Pad = lipgloss.NewStyle().Padding(0, 14)

var Shrtct = style.Info.
	Foreground(lipgloss.Color(theme.Colors.Tertiary)).
	Faint(false)

var BoxStyle = style.BoxStyle.
	Width(shared.Term.Width).
	Height(shared.Term.Height)
