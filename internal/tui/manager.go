package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rrajvardhan/bunkr/internal/server"
	"github.com/rrajvardhan/bunkr/internal/tui/router"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

type Manager struct {
	current tea.Model
}

func InitManager(model tea.Model) Manager {
	return Manager{
		current: model,
	}
}

func (m Manager) Init() tea.Cmd {
	return m.current.Init()
}

func (m Manager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch nav := msg.(type) {
	case shared.NavigateMsg:
		m.current = router.Load(nav.Target)
		return m, m.current.Init()

	case shared.ResetTo:
		if shared.SState.Running {
			shared.SState.Password = ""
			server.Quit()
		}
		m.current = router.Load(nav.Target)
		return m, m.current.Init()

	case shared.SetPasswordMsg:
		if !shared.SState.Running {
			shared.SState.Password = nav.Password
			server.Start(":8080", &shared.SState)
		}
		m.current = router.Load(shared.Dashboard)
		return m, m.current.Init()

	}
	newModel, cmd := m.current.Update(msg)
	m.current = newModel
	return m, cmd
}

func (m Manager) View() string {
	return m.current.View()
}

func (m Manager) Run() {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Shutting down BUNKR...")
	fmt.Printf("%v\n\nServer was running at %s\n", shared.SState.Err, shared.SState.URL)

	os.Exit(0)
}
