package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/rrajvardhan/bunkr/internal/tui/dashboard"
	"github.com/rrajvardhan/bunkr/internal/tui/dashboard/files"
	"github.com/rrajvardhan/bunkr/internal/tui/dashboard/upload"
	"github.com/rrajvardhan/bunkr/internal/tui/home"
	"github.com/rrajvardhan/bunkr/internal/tui/home/about"
	"github.com/rrajvardhan/bunkr/internal/tui/home/client"
	"github.com/rrajvardhan/bunkr/internal/tui/home/client/login"
	"github.com/rrajvardhan/bunkr/internal/tui/home/host"
	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

func Load(route shared.View) tea.Model {
	switch route {

	case shared.Home:
		return home.Start()
	case shared.Host:
		return host.Start()
	case shared.About:
		return about.Start()

	case shared.Dashboard:
		return dashboard.Start()
	case shared.Upload:
		return upload.Start()
	case shared.Files:
		return files.Start()

	case shared.Connect:
		return login.Start()
	case shared.Client:
		return client.Start()

	default:
		return home.Start()
	}
}
