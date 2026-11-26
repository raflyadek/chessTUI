package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

//root view

func (m RootModel) View() tea.View {
	var s string

	switch m.state {
	case menuView:
		s = m.menu.View()
	case gameView:
		s = m.game.View()
	case settingView:

	}
	v := tea.NewView(s)
	v.AltScreen = true
	v.BackgroundColor = style.Black
	// v.BackgroundColor = color.RGBA{R: 17, G: 17, B: 132, A: 1}
	return v
}
