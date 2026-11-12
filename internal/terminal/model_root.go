package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/menu"
)

type stateView int

const (
	menuView stateView = iota
	gameView
)

// root model
type RootModel struct {
	state stateView
	//gameView game.GameModel
	menu menu.MenuModel
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func InitRootModel() RootModel {
	return RootModel{
		menu: menu.InitMenuModel(),
	}
}
