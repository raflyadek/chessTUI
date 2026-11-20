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
		case "backspace", "esc":
			return m, func() tea.Msg {
				return BackToMenu{}
			}
		case "enter":
			if m.state.PlayerMove() == "White" {
				m.player = whitePlayer
			} else {
				m.player = blackPlayer
			}
			// if m.state.MoveCounter%2 == 0 {
			// 	m.player = whitePlayer
			// } else {
			// 	m.player = blackPlayer
			// }
		}
	case timer.TickMsg:
		switch m.player {
		case whitePlayer:
			m.whiteTimer, cmd = m.whiteTimer.Update(msg)
		case blackPlayer:
			m.blackTimer, cmd = m.blackTimer.Update(msg)
		}
	case timer.TimeoutMsg:
		m.state.IsTimeOut = true

	}
	return m, cmd
}

// back to menu
type BackToMenu struct{}
