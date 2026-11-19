package game

import (
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/engine"
)

// delegate view
type activePlayer int

const (
	whitePlayer activePlayer = iota
	blackPlayer
)

type GameModel struct {
	board      [8][8]string
	blackTimer timer.Model
	whiteTimer timer.Model
	player     activePlayer
	//move not using piecesMove anymore
	// from  [2]string
	// to    [2]string
	state  *engine.GameState
	width  int
	height int
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
