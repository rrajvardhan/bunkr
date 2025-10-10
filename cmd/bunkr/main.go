package main

import (
	"github.com/rrajvardhan/bunkr/internal/tui"
	"github.com/rrajvardhan/bunkr/internal/tui/dashboard"
)

func main() {
	manager := tui.InitManager(dashboard.Start())
	manager.Run()
}
