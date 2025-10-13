package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/server"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

type Overseer struct {
	current tea.Model
}

func InitOverseer(model tea.Model) Overseer {
	return Overseer{
		current: model,
	}
}

func (m Overseer) Init() tea.Cmd {
	return m.current.Init()
}

func (m Overseer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch nav := msg.(type) {
	case shared.NavigateMsg:
		m.current = Load(nav.Target)
		return m, m.current.Init()

	case shared.ResetTo:
		if shared.SState.Running {
			shared.SState.Password = ""
			server.Quit()
		}
		m.current = Load(nav.Target)
		return m, m.current.Init()

	case shared.SetPasswordMsg:
		if !shared.SState.Running {
			shared.SState.Password = nav.Password
			server.Start(":8080", &shared.SState)
		}
		m.current = Load(shared.Dashboard)
		return m, m.current.Init()

	case shared.SetServerMsg:
		shared.SState.URL = nav.Url
		shared.SState.Password = nav.Password

		m.current = Load(shared.Client)
		return m, m.current.Init()

	}
	newModel, cmd := m.current.Update(msg)
	m.current = newModel
	return m, cmd
}

func (m Overseer) View() string {
	return m.current.View()
}

func (m Overseer) Run() {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Shutting down BUNKR...")
	fmt.Printf("%v\n\nServer was running at %s\n", shared.SState.Err, shared.SState.URL)

	os.Exit(0)
}
