package router

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rrajvardhan/bunkr/internal/tui/dashboard"
	"github.com/rrajvardhan/bunkr/internal/tui/dashboard/upload"
	"github.com/rrajvardhan/bunkr/internal/tui/home"
	"github.com/rrajvardhan/bunkr/internal/tui/home/about"
	"github.com/rrajvardhan/bunkr/internal/tui/home/auth"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func Load(route shared.Route) tea.Model {
	switch route {
	case shared.Home:
		return home.Start()
	case shared.Auth:
		return auth.Start()
	case shared.About:
		return about.Start()

	case shared.Dashboard:
		return dashboard.Start()
	case shared.Upload:
		return upload.Start()
	default:
		return home.Start()
	}
}
