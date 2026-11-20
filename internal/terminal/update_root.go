package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
	"github.com/raflyadek/chessTUI/internal/terminal/menu"
)

//root update

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case game.BackToMenu:
		m.state = menuView
		return m, nil
	case menu.PlayGame:
		//reset game model with every open game mnu
		m.game = game.InitGameModel()
		m.state = gameView
		return m, m.game.Init()
	}
	//delegate msg from active state
	switch m.state {
	case menuView:
		m.menu, cmd = m.menu.Update(msg)
	case gameView:
		m.game, cmd = m.game.Update(msg)
	}

	return m, cmd
}
