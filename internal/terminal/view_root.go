package terminal

import tea "charm.land/bubbletea/v2"

//root view

func (m RootModel) View() tea.View {
	var s string

	switch m.state {
	case menuView:
		s = m.menu.View()
		// case gameView:
		// 	s = m.game.View()
	}

	return tea.NewView(s)
}
