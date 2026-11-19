package menu

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

//delegate view

// view is if m.selected == 1 then y view or j view
func (m MenuModel) View() string {
	var s strings.Builder
	s.WriteString("test\n\n")
	for i, v := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s.WriteString(fmt.Sprintf("%s [%s] \n", cursor, v))
	}
	s.WriteString("\nbawah")
	lipgloss.NewStyle()
	return s.String()
}
