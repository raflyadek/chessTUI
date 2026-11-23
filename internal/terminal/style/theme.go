package style

import (
	"charm.land/lipgloss/v2"
)

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

var (
	// General UI colors
	Black  = lipgloss.Color("#080808")
	White  = lipgloss.Color("#F2F2F2")
	Grey   = lipgloss.Color("#858585")
	Border = lipgloss.Color("#A0A0A0")

	// Chessboard colors
	LightSquare = lipgloss.Color("#D8C4A2")
	DarkSquare  = lipgloss.Color("#70543B")
)

var (
	AppStyle = lipgloss.NewStyle().
			Foreground(White).
			Background(Black)

	HeaderStyle = lipgloss.NewStyle().
			Foreground(White).
			Bold(true)

	SecondaryStyle = lipgloss.NewStyle().
			Foreground(Grey)

	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(Border).
			Padding(1)

	SectionTitleStyle = lipgloss.NewStyle().
				Foreground(White).
				Bold(true).
				MarginBottom(1)

	InputStyle = lipgloss.NewStyle().
			Foreground(White).
			Border(lipgloss.NormalBorder()).
			BorderForeground(Border).
			Padding(0, 1)

	ActiveClockStyle = lipgloss.NewStyle().
				Foreground(White).
				Bold(true)

	InactiveClockStyle = lipgloss.NewStyle().
				Foreground(Grey)
)

// Chessboard cells
var (
	LightSquareStyle = lipgloss.NewStyle().
				Background(LightSquare).
				Foreground(Black).
				Width(5).
				Height(3).
				Align(lipgloss.Center, lipgloss.Center)

	DarkSquareStyle = lipgloss.NewStyle().
			Background(DarkSquare).
			Foreground(White).
			Width(5).
			Height(3).
			Align(lipgloss.Center, lipgloss.Center)
)
