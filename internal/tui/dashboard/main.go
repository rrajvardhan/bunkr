package dashboard

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct {
	cursor   int
	selected shared.View
	options  []shared.View
	url      string
	showQR   bool
}

func Start() State {
	return State{
		cursor:   0,
		selected: shared.None,
		options:  []shared.View{shared.Files, shared.Upload},
		url:      shared.SState.URL,
		showQR:   false,
	}
}

func (m State) Init() tea.Cmd { return nil }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.showQR {
				m.showQR = false
				return m, nil
			}
			return m, func() tea.Msg { return shared.ResetTo{Target: shared.Home} }
		case "q":
			m.showQR = !m.showQR
			return m, nil
		case "tab":
			m.cursor = (m.cursor + 1) % len(m.options)
			m.selected = m.options[m.cursor]
		case "enter", " ":
			m.selected = m.options[m.cursor]
			return m, func() tea.Msg { return shared.NavigateMsg{Target: m.selected} }
		}
	}
	return m, nil
}

func (m State) View() string {
	header := style.Subtle.Render(shared.SState.URL) + "             " + Shrtct.Render("press 'q' to view QR code.")

	var menu string
	for i, opt := range m.options {
		if i == m.cursor {
			menu += style.ActiveChoice.Render(fmt.Sprintf("> %s             ", opt))
		} else {
			menu += fmt.Sprintf("  %s             ", opt)
		}
	}

	body := Pad.Render(menu)

	controls := Info.Render("[Tab] Cycle options  [Enter] Select  [Esc] Back To Home (kill server)")
	divider := style.Divider.Render(strings.Repeat("─", 74))

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s\n%s",
		header,
		body,
		divider,
		controls,
	)

	if m.showQR {
		content = GenerateCode(shared.SState.URL + "/ping")
		return style.TermCenter(content)
	}

	return style.TermCenter(style.BoxStyle.Render(content))
}
