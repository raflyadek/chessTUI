package game

import (
	"charm.land/bubbles/v2/timer"
	tea "charm.land/bubbletea/v2"
)

//update delegate

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	//from to, move generator
	var cmd tea.Cmd
	switch msg := msg.(type) {
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
		case "enter":
			//TODO: create state where its a valid move before update the timer
			move := m.moveInput.Value()
			m.input = move
			board, err := m.state.PlayGameTUI(m.board, move)
			if err != nil {
				m.err = err.Error()
				m.moveInput.Reset()
				return m, nil
			}
			m.err = ""
			m.board = board
			// m.from = m.board[6][4]
			// m.to = m.board[5][4]
			m.moveInput.Reset()
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
