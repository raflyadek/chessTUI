package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal"
)

var (
	host = "0.0.0.0"
	port = "3000"
)

// TODO:
// create a room, player and gameMiddleware when user choose
// "play with a friend" menu and just use the view on game menu later
func main() {
	//ssh
	p := tea.NewProgram(terminal.InitRootModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
}
