package main

import (
	"github.com/rrajvardhan/bunkr/internal/tui"
	"github.com/rrajvardhan/bunkr/internal/tui/home"
)

func main() {
	manager := tui.InitManager(home.Start())
	manager.Run()
}
