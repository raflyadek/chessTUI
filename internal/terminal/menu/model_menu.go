package menu

import (
	tea "charm.land/bubbletea/v2"
)

//delegate model

type MenuModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	isQuit   bool
}

// init
func InitMenuModel() MenuModel {
	return MenuModel{
		choices:  []string{"Login Lichess Account", "Play With Friends", "Quit"},
		selected: make(map[int]struct{}),
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}
