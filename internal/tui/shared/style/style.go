package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/theme"
)

var BoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color(theme.Colors.SelectionsNoAlpha)).
	Width(80).
	Padding(1, 3)

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
