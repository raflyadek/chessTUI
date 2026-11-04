package main

import (
	"fmt"
	"os"

	// "charm.land/bubbles/v2/cursor"
	// "charm.land/bubbles/v2/textarea"
	// "charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	// "charm.land/lipgloss/v2"
)

func main() {
	p := tea.NewProgram(model{true})
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
}

type model struct {
	quit bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	//when its press key
	case tea.KeyMsg:
		if m.quit {
			switch {
			case msg.String() == "y":
				return m, tea.Quit
			default:
				m.quit = false
			}
		}
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	if m.quit {
		return tea.NewView("Qitting? (y/n)")
	} else {
		return tea.NewView("running....")
	}
}
