package guest

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
)

func (m BoardOnlyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// If GameModel emits BackToMenu on LAN, quit the guest session instead of opening the menu.
	if _, ok := msg.(game.BackToMenu); ok {
		return m, tea.Quit
	}
	var cmd tea.Cmd
	m.game, cmd = m.game.Update(msg)
	return m, cmd
}
