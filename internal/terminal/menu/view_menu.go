package menu

import (
	"charm.land/lipgloss/v2"
	"fmt"
	"github.com/raflyadek/chessTUI/internal/style"
	"strings"
)

//delegate view

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
		AccentHighlights.Render("k"),
		AccentHighlights.Render("j"),
		PrimaryText.Render("Navigate"),
		AccentHighlights.Render("Enter"),
		PrimaryText.Render("Select"),
		AccentHighlights.Render("q"),
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
			choices = append(choices, SelectedStyle.Render("> "+v))
		} else {
			choices = append(choices, NormalStyle.Render(" "+v))
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
