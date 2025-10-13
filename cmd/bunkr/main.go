package main

import (
	"github.com/rrajvardhan/bunkr/internal/tui"
	"github.com/rrajvardhan/bunkr/internal/tui/home"
)

func main() {
	overseer := tui.InitOverseer(home.Start())
	overseer.Run()
}
