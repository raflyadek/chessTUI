package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
	"github.com/raflyadek/chessTUI/internal/terminal/menu"
)

//root update

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		var gameCmd, menuCmd tea.Cmd
		m.game, gameCmd = m.game.Update(msg)
		m.menu, menuCmd = m.menu.Update(msg)
		return m, tea.Batch(gameCmd, menuCmd)
	case game.BackToMenu:
		m.state = menuView
		return m, nil
	case menu.PlayGame:
		//reset game model with every open game mnu
		m.game = game.InitGameModel()
		m.state = gameView
		initCmd := m.game.Init()
		if m.width > 0 && m.height > 0 {

			var sizeCmd tea.Cmd
			m.game, sizeCmd = m.game.Update(tea.WindowSizeMsg{
				Width:  m.width,
				Height: m.height,
			})
			return m, tea.Batch(initCmd, sizeCmd)
		}
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
