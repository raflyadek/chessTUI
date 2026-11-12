package game

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/engine"
)

//delegate view

type GameModel struct {
	board [8][8]string
	//move not using piecesMove anymore
	// from  [2]string
	// to    [2]string
	state *engine.GameState
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func InitGameModel() GameModel {
	return GameModel{
		board: engine.InitBoard(),
		state: engine.NewGameState(),
	}
}
