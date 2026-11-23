package style

import "charm.land/lipgloss/v2"

var (
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B08B5B")).
			Background(lipgloss.Color("#493627")).
			Bold(true).
			Padding(0, 1)

	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B08B5B")).
			Padding(0, 1)

	PrimaryText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E7D5B5"))

	SecondaryText = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8C7964"))

	AccentHighlights = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#B08B5B"))
)
