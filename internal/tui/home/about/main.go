package about

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct{}

func Start() State            { return State{} }
func (m State) Init() tea.Cmd { return tea.EnterAltScreen }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		default:
			return m, func() tea.Msg { return shared.NavigateMsg{Target: shared.Home} }
		}
	}
	return m, nil
}

func (m State) View() string {
	header := Header.Render("About")
	tagline := style.Subtle.Render("whfdi enfiwen wio iefhwih")

	body := ("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	controls := Info.Render("[Any Key] Back ")

	divider := style.Divider.Render(strings.Repeat("─", 74))

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s\n%s",
		header,
		tagline,
		body,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
