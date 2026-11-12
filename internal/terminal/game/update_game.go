package game

import tea "charm.land/bubbletea/v2"

//update delegate

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	//from to, move generator
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace", "esc":
			return m, func() tea.Msg {
				return BackToMenu{}
			}
		}
	}
	return m, nil
}

//back to menu
type BackToMenu struct{}
