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

func InitManager(initialScreen tea.Model) Manager {
	return Manager{
		current: initialScreen,
	}
}

func (m Manager) Init() tea.Cmd {
	return m.current.Init()
}

func (m Manager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch nav := msg.(type) {
	case shared.NavigateMsg:
		if nav.Target == shared.Host {
			m.current = router.Load(shared.Auth)
		} else {
			m.current = router.Load(nav.Target)
		}

	case shared.SetPasswordMsg:
		if !shared.SState.Running {
			shared.SState.Running = true

			go func() {
				shared.SState.Password = nav.Password
				server.Start(":8080")
			}()
		}
		m.current = router.Load(shared.Dashboard)
	}
	newModel, cmd := m.current.Update(msg)
	m.current = newModel
	return m, cmd
}

func (m Manager) View() string {
	return m.current.View()
}

func (m *Manager) Run() {
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Shutting down BUNKR...")
	os.Exit(0)
}
