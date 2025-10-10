package dashboard

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct {
	choices  []shared.Route
	cursor   int
	selected shared.Route
}

func Start() State {
	return State{
		choices:  []shared.Route{shared.Host},
		cursor:   0,
		selected: "",
	}
}

func (m State) Init() tea.Cmd { return tea.EnterAltScreen }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.cursor = (m.cursor + 1) % len(m.choices)
		case "enter", " ":
			m.selected = m.choices[m.cursor]
			return m, func() tea.Msg {
				return shared.NavigateMsg{Target: m.selected}
			}
		}

	case tea.WindowSizeMsg:
		shared.Term.Width = msg.Width
		shared.Term.Height = msg.Height
		return m, nil
	}
	return m, nil
}

func (m State) View() string {
	// Hackerish ASCII-style header
	header := HeaderStyle.Render(`
 ▄▄▄▄▄▄    ▄▄    ▄▄  ▄▄▄   ▄▄  ▄▄   ▄▄▄  ▄▄▄▄▄▄   
 ██▀▀▀▀██  ██    ██  ███   ██  ██  ██▀   ██▀▀▀▀██ 
 ██    ██  ██    ██  ██▀█  ██  ██▄██     ██    ██ 
 ███████   ██    ██  ██ ██ ██  █████     ███████  
 ██    ██  ██    ██  ██  █▄██  ██  ██▄   ██  ▀██▄ 
 ██▄▄▄▄██  ▀██▄▄██▀  ██   ███  ██   ██▄  ██    ██ 
 ▀▀▀▀▀▀▀     ▀▀▀▀    ▀▀   ▀▀▀  ▀▀    ▀▀  ▀▀    ▀▀▀
`)

	divider := DividerStyle.Render(strings.Repeat("─", 74))

	var choices strings.Builder
	for i, choice := range m.choices {
		if i == m.cursor {
			choices.WriteString(fmt.Sprintf("› %s\n", ActiveChoice.Render(string(choice))))
		} else {
			choices.WriteString(fmt.Sprintf("  %s\n", InactiveChoice.Render(string(choice))))
		}
	}

	controls := Info.Render(
		"[Tab] Cycle selection   [Enter] Confirm   [Esc] Quit",
	)

	tagline := Subtle.Render(strings.Join([]string{
		"Lorem ipsum dolor sit amet, consectetur",
		"adipiscing elit. Sed do eiusmod tempor",
	}, "\n"))

	content := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n\n%s",
		header,
		choices.String(),
		controls,
		divider,
		tagline,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
