package game

import (
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
)

//update delegate

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	//from to, move generator
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case middleware.MoveUpdatedMsg:
		m.board = msg.Board
		m.currentTurn = msg.Next
		m.historyMove = msg.History
		m.err = ""
		return m, waitForRoomEvent(m.player2.Events)

	case middleware.OpponentDisconnectedMsg:
		m.opponentOut = true
		m.err = "Opponent disconnected from the match"
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "q":
			return m, func() tea.Msg {
				return BackToMenu{}
			}
		case "ctrl+f":
			m.flipped = !m.flipped
			return m, nil
		case "enter":
			move := m.moveInput.Value()
			m.moveInput.Reset()

			// Online/local SSH multiplayer path
			if m.online && m.room != nil {
				if m.opponentOut {
					m.err = "Match ended: opponent disconnected"
					return m, nil
				}
				if err := m.room.SubmitMove(m.mySide, move); err != nil {
					m.err = err.Error()
				}
				cmd = m.switchPlayer()
				return m, cmd
			}

			// local play path
			board, err := m.state.PlayGameTUI(m.board, move)
			if err != nil {
				m.err = err.Error()
				m.moveInput.Reset()
				return m, nil
			}
			//add history
			history := m.state.HistoryMoveHelper(move, m.board)
			m.historyMove = append(m.historyMove, history)
			// if m.state.IsCheckMate {
			// 	m.checkmatee = "checkmateee"
			// } else {
			// 	m.checkmatee = "not checkmate??"
			// }

			m.err = ""
			m.board = board
			cmd = m.switchPlayer()
			return m, cmd
		default:
			m.moveInput, cmd = m.moveInput.Update(msg)
		}
	case timer.TickMsg:

		if msg.ID == m.whiteTimer.ID() && m.player == whitePlayer {
			m.whiteTimer, cmd = m.whiteTimer.Update(msg)
		} else if msg.ID == m.blackTimer.ID() && m.player == blackPlayer {
			m.blackTimer, cmd = m.blackTimer.Update(msg)
		} // cmd = tea.Batch(whiteCmd, blackCmd)
	case timer.TimeoutMsg:
		m.state.IsTimeOut = true

	}
	return m, cmd
}

// helper switch timer
func (m *GameModel) switchPlayer() tea.Cmd {
	switch m.player {
	case whitePlayer:
		m.player = blackPlayer
		// m.blackTimer.Stop()
		return m.blackTimer.Init()

	case blackPlayer:
		m.player = whitePlayer
		return m.whiteTimer.Init()
	}

	return nil
}

// back to menu
type BackToMenu struct{}

func (m *GameModel) GetBoard() [8][8]string {
	return m.state.InitBoard()
}

func (m *GameModel) GetPlayGame() func([8][8]string, string) ([8][8]string, error) {
	return m.state.PlayGameTUI
}
