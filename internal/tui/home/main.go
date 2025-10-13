package home

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
		choices:  []shared.Route{shared.Host, shared.Connect, shared.About},
		cursor:   0,
		selected: shared.None,
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
			return m, func() tea.Msg { return shared.NavigateMsg{Target: m.selected} }
		}

	case tea.WindowSizeMsg:
		shared.Term.Width = msg.Width
		shared.Term.Height = msg.Height
		return m, nil
	}
	return m, nil
}

func (m State) View() string {
	header := style.Header.Render(`
 ▄▄▄▄▄▄    ▄▄    ▄▄  ▄▄▄   ▄▄  ▄▄   ▄▄▄  ▄▄▄▄▄▄   
 ██▀▀▀▀██  ██    ██  ███   ██  ██  ██▀   ██▀▀▀▀██ 
 ██    ██  ██    ██  ██▀█  ██  ██▄██     ██    ██ 
 ███████   ██    ██  ██ ██ ██  █████     ███████  
 ██    ██  ██    ██  ██  █▄██  ██  ██▄   ██  ▀██▄ 
 ██▄▄▄▄██  ▀██▄▄██▀  ██   ███  ██   ██▄  ██    ██ 
 ▀▀▀▀▀▀▀     ▀▀▀▀    ▀▀   ▀▀▀  ▀▀    ▀▀  ▀▀    ▀▀▀
`)

	divider := style.Divider.Render(strings.Repeat("─", 74))

	var choices strings.Builder
	for i, choice := range m.choices {
		if i == m.cursor {
			choices.WriteString(style.ActiveChoice.Render(fmt.Sprintf("> %s             ", string(choice))))
		} else {
			choices.WriteString(fmt.Sprintf("  %s             ", (string(choice))))
		}
	}

	controls := style.Info.Render(
		"[Tab] Cycle selection   [Enter] Confirm   [Esc] Quit",
	)

	tagline := style.Subtle.Render("Share files locally without trusting the cloud, your ISP, or that one guy who definitely logs everything.")

	modes := Pad.Render(choices.String())

	content := fmt.Sprintf(
		"%s\n%s\n\n%s\n\n%s\n%s",
		header,
		tagline,
		modes,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
