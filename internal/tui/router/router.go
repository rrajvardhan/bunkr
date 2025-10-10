package router

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rrajvardhan/bunkr/internal/tui/auth"
	"github.com/rrajvardhan/bunkr/internal/tui/dashboard"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func Load(route shared.Route) tea.Model {
	switch route {
	case shared.Dashboard:
		return dashboard.Start()
	case shared.Auth:
		return auth.Start()
	default:
		return dashboard.Start()
	}
}
