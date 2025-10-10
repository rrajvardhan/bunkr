package auth

import (
	"fmt"

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
			return m, tea.Quit

		case tea.KeyTab:
			m.cursor = (m.cursor + 1) % 2
		case tea.KeyBackspace, tea.KeyDelete:

			m.errorMessage = ""
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

			m.errorMessage = ""
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

	passLine := fmt.Sprintf("Password: %s", hide(m.password))
	confLine := fmt.Sprintf("Confirm:  %s", hide(m.confirm))
	if m.cursor == 0 {
		passLine = Cursor.Render(passLine)
	} else {
		confLine = Cursor.Render(confLine)
	}

	body := passLine + "\n" + confLine
	if m.errorMessage != " " {
		body += "\n\n" + Error.Render(m.errorMessage)
	}

	controls := Info.Render(
		"[Enter]  Confirm            [Tab] Cycle field\n[Ctrl+H] Toggle visibility  [Esc] Quit",
	)

	content := fmt.Sprintf(
		"\n%s\n\n%s",
		body,
		controls,
	)

	return style.TermCenter(style.BoxStyle.Render(content))
}
