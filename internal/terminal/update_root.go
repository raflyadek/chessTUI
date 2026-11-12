package terminal

import tea "charm.land/bubbletea/v2"

//root update

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	//delegate msg from active state
	switch m.state {
	case menuView:
		m.menu, cmd = m.menu.Update(msg)
	}

	return m, cmd
}
