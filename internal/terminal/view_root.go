package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

func (m RootModel) View() tea.View {
	var s string

	switch m.state {
	case menuView:
		s = m.menu.View()
	case gameView:
		s = m.game.View()
	case settingView:
		s = m.setting.View()
	}

	v := tea.NewView(s)
	v.AltScreen = true
	v.BackgroundColor = style.Black
	return v
}
