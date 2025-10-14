package about

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

var line = "Private file sharing for local networks."

var desc = ("Bunkr is a lightweight, LAN-only file sharing tool built in Go. " +
	"It encrypts every file before transfer and never touches the cloud. " +
	"Designed for simplicity, speed, and a bit of paranoia. " +
	"\n\n" +
	"The terminal interface is powered by Bubble Tea.\n" +
	"You can find the source code at github.com/rrajvardhan/bunkr.")

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

	tagline := style.Subtle.Render(line)
	body := desc

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
