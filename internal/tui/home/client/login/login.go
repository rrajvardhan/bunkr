package login

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
	"github.com/rrajvardhan/bunkr/internal/tui/shared/style"
)

type State struct {
	serverURL    string
	password     string
	cursor       int  // 0 = URL, 1 = password
	hidden       bool // hide password input
	errorMessage string
}

func Start() State {
	return State{
		serverURL: "",
		password:  "",
		cursor:    0,
		hidden:    true,
	}
}

func (m State) Init() tea.Cmd { return tea.EnterAltScreen }

func (m State) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyTab:
			m.cursor = (m.cursor + 1) % 2

		case tea.KeyBackspace, tea.KeyDelete:
			m.errorMessage = ""
			if m.cursor == 0 && len(m.serverURL) > 0 {
				m.serverURL = m.serverURL[:len(m.serverURL)-1]
			} else if m.cursor == 1 && len(m.password) > 0 {
				m.password = m.password[:len(m.password)-1]
			}

		case tea.KeyCtrlH:
			m.hidden = !m.hidden

		case tea.KeyEnter:
			if strings.TrimSpace(m.serverURL) == "" {
				m.errorMessage = "Server URL cannot be empty!"
				return m, nil
			}

			if err := AuthenticateWithServer(m.serverURL, m.password); err != nil {
				m.errorMessage = fmt.Sprintf("Failed to connect: %v", err)
				return m, nil
			}

			return m, func() tea.Msg {
				return shared.SetServerMsg{
					Url:      m.serverURL,
					Password: m.password,
				}
			}

		default:

			if len(msg.String()) == 1 {
				if m.cursor == 0 {
					m.serverURL += msg.String()
				} else {
					m.password += msg.String()
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

	header := Header.Render("Connect to Server")
	info := style.Subtle.Render("Enter the server URL and password (if required) to connect.\nLeave password blank for no authentication.\n")

	urlLine := fmt.Sprintf("  Server URL : %s ", m.serverURL)
	passLine := fmt.Sprintf("  Password   : %s ", hide(m.password))

	if m.cursor == 0 {
		urlLine = style.ActiveChoice.Render(urlLine + "<")
	} else {
		passLine = style.ActiveChoice.Render(passLine + "<")
	}

	body := urlLine + "\n" + passLine
	body += "\n\n" + Error.Render(m.errorMessage)

	controls := Info.Render("[Enter] Connect  [Tab] Switch Field  [Ctrl+H] Toggle Password  [Esc] Back")
	divider := style.Divider.Render(strings.Repeat("─", 74))

	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n%s", header, info, body, divider, controls)

	return style.TermCenter(style.BoxStyle.Render(content))
}
