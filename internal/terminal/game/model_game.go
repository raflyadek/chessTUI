package game

import (
	"time"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/engine"
	"github.com/raflyadek/chessTUI/internal/middleware"
)

// delegate view
type activePlayer int

const defaultTime = 1 * time.Minute
const (
	whitePlayer activePlayer = iota
	blackPlayer
)

type GameModel struct {
	board       [8][8]string
	blackTimer  timer.Model
	whiteTimer  timer.Model
	player      activePlayer
	moveInput   textinput.Model
	historyMove []string
	flipped     bool
	isLocal     bool
	isLan       bool
	state       *engine.GameState
	width       int
	height      int
	err         string
	// Add these fields to your existing GameModel struct:
	room        *middleware.Room
	player2     *middleware.Player
	mySide      middleware.Side
	currentTurn middleware.Side
	online      bool
	opponentOut bool
}

func (m GameModel) Init() tea.Cmd {
	if m.online && m.player2 != nil {
		return tea.Batch(
			m.whiteTimer.Init(),
			waitForRoomEvent(m.player2.Events),
		)
	}
	return m.whiteTimer.Init()
}
func waitForRoomEvent(ch <-chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func InitGameModel() GameModel {
	input := textinput.New()
	input.Placeholder = "Enter move (e.g. e2e4)"
	input.SetWidth(30)
	input.Focus()
	return GameModel{
		board:      engine.InitBoard(),
		state:      engine.NewGameState(),
		whiteTimer: timer.New(defaultTime),
		blackTimer: timer.New(defaultTime),
		moveInput:  input,
	}
}

// InitOnlineGameModel initializes GameModel for an SSH LAN match
func InitOnlineGameModel(
	room *middleware.Room,
	player *middleware.Player,
	assigned middleware.Side,
	tf middleware.TimeFormat,
	board [8][8]string,
) GameModel {
	m := InitGameModel()
	m.room = room
	m.player2 = player
	m.mySide = assigned
	m.currentTurn = middleware.SideWhite
	m.online = true
	m.board = board
	m.flipped = (assigned == middleware.SideBlack) // Auto-flip for Black!

	// Set clocks from the room's chosen TimeFormat
	if tf.Initial > 0 {
		m.whiteTimer = timer.New(tf.Initial)
		m.blackTimer = timer.New(tf.Initial)
	}
	return m
}
