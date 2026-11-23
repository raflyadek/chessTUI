package game

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

type Layout struct {
	BoardWidth   int
	BoardHeight  int
	SidebarWidth int
	TotalWidth   int
	TotalHeight  int
}

func (m GameModel) calculateLayout() Layout {
	boardWidth := 48  // adjust to your actual rendered board width
	boardHeight := 30 // adjust to your actual rendered board height
	sidebarWidth := 26

	return Layout{
		BoardWidth:   boardWidth,
		BoardHeight:  boardHeight,
		SidebarWidth: sidebarWidth,
		TotalWidth:   boardWidth + sidebarWidth + 2,
		TotalHeight:  boardHeight + 10,
	}
}

// delegate view

func (m GameModel) View() string {
	// layout := m.calculateLayout()

	board := m.boardRender()
	boardWidth := lipgloss.Width(board)
	//TODO:if its flipped

	blackTimer := m.clockRender(
		blackPlayer,
		m.blackTimerRender(),
		boardWidth,
	)
	whiteTimer := m.clockRender(
		whitePlayer,
		m.whiteTimerRender(),
		boardWidth,
	)
	sidebar := m.sidebarRender()
	footer := m.footerRender()
	boardColumn := lipgloss.JoinVertical(
		lipgloss.Left,
		blackTimer,
		board,
		whiteTimer,
	)
	main := lipgloss.JoinHorizontal(
		lipgloss.Top,
		boardColumn,
		sidebar,
	)
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		main,
		footer,
	)

	// return lipgloss.JoinVertical(
	// 	lipgloss.Left,
	// 	main,
	// 	footer,
	// )
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
	// moveInput := m.InputRender()
	// player := m.state.PlayerMove()

	// playerState := m.player
	// player := ""
	// if playerState == whitePlayer {
	// 	player = "whitE"
	// } else {
	// 	player = "Black"
	// }
	// s := m.err

	// from := "from " + m.from + "|"
	// to := "to " + m.to + "|"
	// input := "input " + m.input
	// checkmate := m.checkmatee
	// history := m.historyRender()
	// return lipgloss.JoinVertical(lipgloss.Top, blackTimer, board, whiteTimer, moveInput, history)
}

// render board
func (m GameModel) boardRender() string {
	var rows []string

	rankStyle := lipgloss.NewStyle().
		Foreground(style.SecondaryStyle.GetForeground()).
		Width(2).
		Align(lipgloss.Right)

	fileStyle := lipgloss.NewStyle().
		Foreground(style.SecondaryStyle.GetForeground()).
		Width(5).
		Align(lipgloss.Center)
	pieces := map[string]string{
		"P": "♟︎",
		"R": "♜",
		"N": "♞",
		"B": "♝",
		"Q": "♛",
		"K": "♚",
		"p": "♙",
		"r": "♖",
		"n": "♘",
		"b": "♗",
		"q": "♕",
		"k": "♔",
	}
	var files []string
	files = append(files, "  ")

	for j := 0; j < 8; j++ {
		files = append(files, fileStyle.Render(
			string(rune('a'+j)),
		))
	}

	rows = append(rows, lipgloss.JoinHorizontal(
		lipgloss.Top, files...,
	))

	for i := 0; i < len(m.board); i++ {
		var cells []string
		cells = append(cells, rankStyle.Render(
			fmt.Sprintf("%d", 8-i),
		))
		for j := 0; j < len(m.board[i]); j++ {
			piece := m.board[i][j]
			content := " "
			if piece != "" {
				content = pieces[piece]
			}
			styles := style.LightSquareStyle
			if (i+j)%2 == 0 {
				styles = style.DarkSquareStyle
			}
			cells = append(cells, styles.Render(content))
		}
		cells = append(cells, rankStyle.Render(
			fmt.Sprintf("%d", 8-i),
		))
		rows = append(rows, lipgloss.JoinHorizontal(
			lipgloss.Top, cells...,
		))

	}
	rows = append(rows, lipgloss.JoinHorizontal(
		lipgloss.Top, files...,
	))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		rows...,
	)
}

func (m GameModel) sidebarRender() string {
	historyTitle := style.SectionTitleStyle.Render("HISTORY MOVE")
	history := m.historyRender()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		historyTitle,
		strings.Repeat("-", 24),
		history,
	)
	return style.PanelStyle.Width(26).Height(30).Render(content)
}

func (m GameModel) footerRender() string {
	input := style.InputStyle.Render(
		m.InputRender(),
	)

	help := style.SecondaryStyle.Render(
		"Enter: Submit\t|\tEsc: Menu\t|\tCtrl+C: Quit",
	)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		input,
		help,
	)
}

// render clock
func (m GameModel) blackTimerRender() string {

	s := m.blackTimer.View()
	return s
}

func (m GameModel) whiteTimerRender() string {
	s := m.whiteTimer.View()
	return s
}

// render move input
func (m GameModel) InputRender() string {
	return m.moveInput.View()
}

// render history
func (m GameModel) historyRender() string {
	var s strings.Builder

	for i := 0; i < len(m.historyMove); i += 2 {
		moveNumber := (i / 2) + 1

		whiteMove := m.historyMove[i]
		blackMove := ""

		if i+1 < len(m.historyMove) {
			blackMove = m.historyMove[i+1]
		}

		fmt.Fprintf(&s, "%d. %-6s %s\n",
			moveNumber,
			whiteMove,
			blackMove,
		)
	}

	return strings.TrimRight(s.String(), "\n")

}

// helper for timer
func (m GameModel) clockRender(
	player activePlayer,
	timer string,
	boardWidth int,
) string {
	indicator := "  "

	if m.player == player {
		indicator = "▶ "
	}

	content := indicator + timer

	return lipgloss.NewStyle().
		Width(boardWidth).
		Align(lipgloss.Right).
		Render(content)
}
