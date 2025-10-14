package about

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

var line = "Encrypted file sharing that stays on your network."

var desc = ("Share files over your local network with password protection and encryption. " +
"For when your devices are on the same network and the cloud is unnecessary. " +
"No external servers needed — just direct, encrypted transfers.\n" + 
    "Built in Go with a Bubble Tea terminal interface." +
    "\n\n" +
    "github.com/rrajvardhan/bunkr")

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
