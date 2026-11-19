package terminal

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

//root view

func (m RootModel) View() tea.View {
	var s string

	switch m.state {
	case menuView:
		s = m.menu.View()
	case gameView:
		s = m.game.View()
	}
	v := tea.NewView(s)
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Color("#17120F")
	// v.BackgroundColor = color.RGBA{R: 17, G: 17, B: 132, A: 1}
	return v
}
