package guest

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

func (m BoardOnlyModel) View() tea.View {
	v := tea.NewView(m.game.View())
	v.AltScreen = true
	v.BackgroundColor = style.Black
	return v
}
