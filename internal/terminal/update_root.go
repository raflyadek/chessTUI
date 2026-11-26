package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
	"github.com/raflyadek/chessTUI/internal/terminal/menu"
	"github.com/raflyadek/chessTUI/internal/terminal/setting"
)

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		var gameCmd, menuCmd, settingCmd tea.Cmd
		m.game, gameCmd = m.game.Update(msg)
		m.menu, menuCmd = m.menu.Update(msg)
		m.setting, settingCmd = m.setting.Update(msg)
		return m, tea.Batch(gameCmd, menuCmd, settingCmd)

	case game.BackToMenu, setting.BackToMenuMsg:
		m.state = menuView
		return m, nil

	// case menu.PlayGame:
	// 	// Local single-player / pass-and-play game
	// 	m.game = game.InitGameModel()
	// 	m.state = gameView
	// 	return m, m.syncGameSize(m.game.Init())

	case menu.SettingGame:
		// Open Play With Friends (Create / Join Room)
		m.state = settingView
		return m, m.setting.Init()

	case setting.StartOnlineGameMsg:
		// Both SSH players matched! Start online GameModel
		m.game = game.InitOnlineGameModel(
			msg.Room,
			msg.Player,
			msg.Assigned,
			msg.TimeFormat,
			msg.Board,
		)
		m.state = gameView
		return m, m.syncGameSize(m.game.Init())
	}

	// Delegate messages to the active screen
	switch m.state {
	case menuView:
		m.menu, cmd = m.menu.Update(msg)
	case gameView:
		m.game, cmd = m.game.Update(msg)
	case settingView:
		m.setting, cmd = m.setting.Update(msg)
	}

	return m, cmd
}

func (m *RootModel) syncGameSize(initCmd tea.Cmd) tea.Cmd {
	if m.width > 0 && m.height > 0 {
		var sizeCmd tea.Cmd
		m.game, sizeCmd = m.game.Update(tea.WindowSizeMsg{
			Width:  m.width,
			Height: m.height,
		})
		return tea.Batch(initCmd, sizeCmd)
	}
	return initCmd
}
