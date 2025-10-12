package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var (
	Header = lipgloss.NewStyle().
		Bold(true).
		Padding(0, 9).
		Foreground(lipgloss.Color(theme.Colors.Primary))

	Divider = lipgloss.NewStyle().
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
		Foreground(lipgloss.Color(theme.Colors.Quinary))

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(theme.Colors.SelectionsNoAlpha)).
			Width(80).
			Padding(1, 3, 0, 3)

	Error = lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Senary)).
		Bold(true)
)

func TermCenter(content string) string {
	return lipgloss.Place(
		shared.Term.Width,
		shared.Term.Height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func HorizontalSplit(left, right string, margin int) string {
	if margin < 0 {
		margin = 0
	}

	leftStyled := lipgloss.NewStyle().
		MarginRight(margin).
		Render(left)

	rightStyled := lipgloss.NewStyle().
		MarginLeft(margin).
		Render(right)

	return lipgloss.JoinHorizontal(lipgloss.Top, leftStyled, rightStyled)
}
