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
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	}
}

type model struct {
	board    [8][8]string
	choices  []string
	cursor   int
	selectde map[int]struct{}
}

func initialModel() model {
	board := [8][8]string{}

	//white
	board[6] = [8]string{"P", "P", "P", "P", "P", "P", "P", "P"}
	board[7] = [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"}

	//black
	board[1] = [8]string{"p", "p", "p", "p", "p", "p", "p", "p"}
	board[0] = [8]string{"r", "n", "b", "q", "k", "b", "n", "r"}
	return model{
		board:    board,
		choices:  []string{"jalan 1 langkah", "keluar"},
		selectde: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", "space":
			_, ok := m.selectde[m.cursor]
			if ok {
				delete(m.selectde, m.cursor)
			} else {
				m.selectde[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	//header
	s := "what mode you want to play?\n\n"

	for row, _ := range m.board {
		for col, piece := range m.board[row] {
			if piece == "" {
				if (col+row)%2 == 0 {
					s += "  .  "
				} else {
					s += "  ,  "
				}
			} else {
				s += fmt.Sprintf("  %s  ", piece)
			}
			if col == 7 {
				s += "\n"
			}
		}
		s += "\n"
	}
	for i, choice := range m.choices {
		cursor := " "
		//if the cursor pointing the choice
		if m.cursor == i {
			cursor = ">"
		}
		//is it selected?
		checked := " "
		if _, ok := m.selectde[i]; ok {
			checked = "x"
		}

		//render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	//footer
	s += "\n\ndont forget to play chess once a day!"
	return tea.NewView(s)
}
