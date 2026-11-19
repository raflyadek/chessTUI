package menu

import (
	"fmt"
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

	PrimaryText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E7D5B5"))

	secondaryText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8C7964"))

	accentHighlights = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B08B5B"))
)

// view is if m.selected == 1 then y view or j view
func (m MenuModel) View() string {
	header := m.headerRender()
	footer := m.footerRender()
	content := m.contentRender(header, footer)
	//its refresh everytime this function called
	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}

func (m MenuModel) headerRender() string {
	headerContent := `
   ██████╗██╗  ██╗███████╗███████╗███████╗
  ██╔════╝██║  ██║██╔════╝██╔════╝██╔════╝
  ██║     ███████║█████╗  ███████╗███████╗
  ██║     ██╔══██║██╔══╝  ╚════██║╚════██║
  ╚██████╗██║  ██║███████╗███████║███████║
   ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
	`
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Bold(true).
		Foreground(lipgloss.Color("#E8C88A")).
		Render(headerContent)
	return header
}

func (m MenuModel) footerRender() string {
	footerContent := fmt.Sprintf("%s %s\t%s\t%s\t%s\t%s\t%s",
		accentHighlights.Render("k"),
		accentHighlights.Render("j"),
		PrimaryText.Render("Navigate"),
		accentHighlights.Render("Enter"),
		PrimaryText.Render("Select"),
		accentHighlights.Render("q"),
		PrimaryText.Render("Quit"))
	footer := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		Render(footerContent)
	return footer
}

func (m MenuModel) contentRender(header, footer string) string {

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

	return content
}
