package menu

import tea "charm.land/bubbletea/v2"

//update delegate
func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
			tea.Sequence()
		case "j", "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			// if m.isQuit {
			// 	return m, tea.Quit
			// }
			if m.cursor == 2 {
				return m, tea.Quit
			}
			if m.cursor == 1 {
				return m, func() tea.Msg {
					return PlayGame{}
				}
			}
		}
	}
	return m, nil
}

type PlayGame struct{}
