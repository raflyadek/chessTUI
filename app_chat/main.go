package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal"
)

func main() {
	p := tea.NewProgram(terminal.InitRootModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
}
