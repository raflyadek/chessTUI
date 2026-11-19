package menu

import (
	"strings"

	"charm.land/lipgloss/v2"
)

//delegate view

var (
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B08B5B")).
			Background(lipgloss.Color("#493627")).
			Bold(true).
			Padding(0, 1)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B08B5B")).
			Padding(0, 1)
)

// view is if m.selected == 1 then y view or j view
func (m MenuModel) View() string {
	headerContent := `
   ██████╗██╗  ██╗███████╗███████╗███████╗
  ██╔════╝██║  ██║██╔════╝██╔════╝██╔════╝
  ██║     ███████║█████╗  ███████╗███████╗
  ██║     ██╔══██║██╔══╝  ╚════██║╚════██║
  ╚██████╗██║  ██║███████╗███████║███████║
   ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
	`
	footerContent := "k j 	Navigate	Enter	Select	q	Quit"
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Bold(true).
		Foreground(lipgloss.Color("#E8C88A")).
		Render(headerContent)
	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Render(footerContent)
	var choices []string
	for i, v := range m.choices {
		if m.cursor == i {
			choices = append(choices, selectedStyle.Render("> "+v))
		} else {
			choices = append(choices, normalStyle.Render(" "+v))
		}
	}
	content := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)).
		Align(lipgloss.Center, lipgloss.Center).
		Bold(true).
		Render(strings.Join(choices, "\n"))
	// return s.String()
	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}
