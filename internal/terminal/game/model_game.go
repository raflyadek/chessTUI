package game

import (
	"time"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/engine"
)

// delegate view
type activePlayer int

const defaultTime = 1 * time.Minute
const (
	whitePlayer activePlayer = iota
	blackPlayer
)

type GameModel struct {
	board      [8][8]string
	blackTimer timer.Model
	whiteTimer timer.Model
	player     activePlayer
	moveInput  textinput.Model
	//move not using piecesMove anymore
	// from  [2]string
	// to    [2]string
	state  *engine.GameState
	width  int
	height int
	err    string
	from   string
	to     string
	input  string
}

func (m GameModel) Init() tea.Cmd {
	return m.whiteTimer.Init()
}

func InitGameModel() GameModel {
	input := textinput.New()
	input.Placeholder = "Pnter move (e.g. e2 e4)"
	input.Focus()
	return GameModel{
		board:      engine.InitBoard(),
		state:      engine.NewGameState(),
		whiteTimer: timer.New(defaultTime),
		blackTimer: timer.New(defaultTime),
		moveInput:  input,
	}
}
