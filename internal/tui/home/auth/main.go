package auth

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct {
	password     string
	confirm      string
	cursor       int
	hidden       bool
	errorMessage string
}

func Start() State {
	return State{
		password: "",
		confirm:  "",
		hidden:   true,
	}
}

func (m State) Init() tea.Cmd { return tea.EnterAltScreen }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:

			return m, func() tea.Msg { return shared.NavigateMsg{Target: shared.Home} }

		case tea.KeyTab:
			m.cursor = (m.cursor + 1) % 2
		case tea.KeyBackspace, tea.KeyDelete:

			m.errorMessage = " "
			if m.cursor == 0 && len(m.password) > 0 {
				m.password = m.password[:len(m.password)-1]
			} else if m.cursor == 1 && len(m.confirm) > 0 {
				m.confirm = m.confirm[:len(m.confirm)-1]
			}

		case tea.KeyEnter:
			if m.password != m.confirm {
				m.errorMessage = "Passwords do not match!"
				return m, nil
			}
			return m, func() tea.Msg { return shared.SetPasswordMsg{Password: m.password} }

		default:

			switch msg.String() {
			case "ctrl+h":
				m.hidden = !m.hidden
			default:
				if len(msg.String()) == 1 {
					if m.cursor == 0 {
						m.password += msg.String()
					} else {
						m.confirm += msg.String()
					}
				}
			}
		}
	}
	return m, nil
}

func (m State) View() string {
	hide := func(s string) string {
		if !m.hidden {
			return s
		}
		stars := ""
		for range s {
			stars += "*"
		}
		return stars
	}

	header := Header.Render("Set a password")
	info := style.Subtle.Render("This password will be used to authenticate devices connecting to this server.\nplease keep it safe.\n\n(Leave the field blank to allow anyone to connect without a password)")

	passLine := fmt.Sprintf("  Password : %s ", hide(m.password))
	confLine := fmt.Sprintf("  Confirm  : %s ", hide(m.confirm))
	if m.cursor == 0 {
		passLine = style.ActiveChoice.Render(passLine + "<")
	} else {
		confLine = style.ActiveChoice.Render(confLine + "<")
	}

	body := passLine + "\n" + confLine
	body += "\n\n" + Error.Render(m.errorMessage)

	controls := Info.Render(
		"[Enter] Confirm  [Tab] Cycle field  [Ctrl+H] Toggle visibility  [Esc] Back",
	)

	divider := style.Divider.Render(strings.Repeat("─", 74))

	content := fmt.Sprintf(
		"%s\n\n%s\n\n\n%s\n\n%s\n%s",
		header,
		info,
		body,
		divider,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
