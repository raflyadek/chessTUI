package game

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

// delegate view

func (m GameModel) View() string {
	//if its still menu
	layout := m.calculateLayout()
	board := m.boardRender(layout.cellWidth, layout.cellHeight)
	boardWidth := lipgloss.Width(board)
	boardHeight := lipgloss.Height(board)

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

	//if flipped / play as a black
	topTimer := blackTimer
	botTimer := whiteTimer
	if m.flipped {
		topTimer = whiteTimer
		botTimer = blackTimer
	}

	footer := m.footerRender(boardWidth)
	boardColumn := lipgloss.JoinVertical(
		lipgloss.Left,
		topTimer,
		board,
		botTimer,
		footer,
	)
	var main string
	if layout.stacked {
		stackSidebar := max(20, boardWidth-2)
		sidebar := m.sidebarRender(stackSidebar, 8)
		main = lipgloss.JoinVertical(
			lipgloss.Left,
			boardColumn,
			sidebar,
		)
	} else {
		sidebar := m.sidebarRender(layout.sidebarWidth, boardHeight)
		main = lipgloss.JoinHorizontal(
			lipgloss.Top,
			boardColumn,
			sidebar,
		)
	}
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		main,
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

// render board
func (m GameModel) boardRender(cellWidth, cellHeight int) string {
	var rows []string

	rankStyle := lipgloss.NewStyle().
		Foreground(style.SecondaryStyle.GetForeground()).
		Width(2).
		Height(cellHeight).
		Align(lipgloss.Right, lipgloss.Center)

	fileStyle := lipgloss.NewStyle().
		Foreground(style.SecondaryStyle.GetForeground()).
		Width(cellWidth).
		Align(lipgloss.Center)
	lightSquare := style.LightSquareStyle.Width(cellWidth).Height(cellHeight)
	darkSquare := style.DarkSquareStyle.Width(cellWidth).Height(cellHeight)

	for i := 0; i < len(m.board); i++ {
		boardRow := i
		if m.flipped {
			boardRow = 7 - i
		}
		var cells []string
		cells = append(cells, rankStyle.Render(
			fmt.Sprintf("%d", 8-boardRow),
		))
		for j := 0; j < len(m.board[i]); j++ {
			boardCol := j
			if m.flipped {
				boardCol = 7 - j
			}
			piece := m.board[boardRow][boardCol]
			sqStyle := lightSquare
			if (boardRow+boardCol)%2 == 0 {
				sqStyle = darkSquare
			}

			content := " "
			if piece != "" {
				pieceRender := m.renderSquareArt(piece, cellWidth, cellHeight, sqStyle)
				content = pieceRender
			}
			cells = append(cells, sqStyle.Render(content))
		}
		rows = append(rows, lipgloss.JoinHorizontal(
			lipgloss.Top, cells...,
		))

	}

	var files []string
	files = append(files, "  ")

	for j := 0; j < 8; j++ {
		boardCol := j
		if m.flipped {
			boardCol = 7 - j
		}
		files = append(files, fileStyle.Render(
			string(rune('a'+boardCol)),
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

func (m GameModel) sidebarRender(width, height int) string {
	historyTitle := style.SectionTitleStyle.Render("HISTORY MOVE")
	dividerWidth := max(8, width-2)
	history := m.historyRender()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		historyTitle,
		strings.Repeat("-", dividerWidth),
		history,
	)
	return style.PanelStyle.Width(width).Height(height).Render(content)
}

func (m GameModel) footerRender(boardWidth int) string {
	input := style.InputStyle.Width(boardWidth).Render(
		m.InputRender(),
	)

	help := style.SecondaryStyle.Width(boardWidth).Render(
		"Enter: Submit  |  Esc: Menu  |  Ctrl+C: Quit",
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

// Medium pieces: 5 chars wide x 3 lines tall (for cellWidth >= 6, cellHeight >= 3)
var mediumPieces = map[string][]string{
	"P": {
		"  ▄  ",
		" ▐█▌ ",
		" ▀▀▀ ",
	},
	"R": {
		" █▀█ ",
		" ▐█▌ ",
		" ▀▀▀ ",
	},
	"N": {
		" ▄█▄ ",
		" ▀█▌ ",
		" ▀▀▀ ",
	},
	"B": {
		"  ◆  ",
		" ▐█▌ ",
		" ▀▀▀ ",
	},
	"Q": {
		" ▼ ▼ ",
		" ▐█▌ ",
		" ▀▀▀ ",
	},
	"K": {
		"  ✚  ",
		" ▐█▌ ",
		" ▀▀▀ ",
	},
}

// Fits inside cellWidth >= 5 and cellHeight >= 3
var compactPieces = map[string][]string{
	"P": {
		" ▄ ",
		"▐█▌",
		" ▀ ",
	},
	"R": {
		"█▀█",
		"▐█▌",
		" ▀ ",
	},
	"N": {
		"▄█▄",
		"▀█▌",
		" ▀ ",
	},
	"B": {
		" ◆ ",
		"▐█▌",
		" ▀ ",
	},
	"Q": {
		"▼▼▼",
		"▐█▌",
		" ▀ ",
	},
	"K": {
		" ✚ ",
		"▐█▌",
		" ▀ ",
	},
}

// Large pieces: 7 chars wide x 4 lines tall (for cellWidth >= 8, cellHeight >= 4)
var largePieces = map[string][]string{
	"P": {
		"       ",
		"  ▄█▄  ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
	"R": {
		" █▄█▄█ ",
		"  ▐█▌  ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
	"N": {
		" ▄██▄  ",
		" ███▌  ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
	"B": {
		"  ▄◆▄  ",
		"  ▐█▌  ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
	"Q": {
		" ▼ ▼ ▼ ",
		" ▐███▌ ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
	"K": {
		"  ▄✚▄  ",
		" ▐███▌ ",
		"  ▐█▌  ",
		" ▄███▄ ",
	},
}

var unicodePieces = map[string]string{
	"P": "♟︎", "R": "♜", "N": "♞", "B": "♝", "Q": "♛", "K": "♚",
	"p": "♙", "r": "♖", "n": "♘", "b": "♗", "q": "♕", "k": "♔",
}

func (m GameModel) renderSquareArt(
	piece string,
	cellWidth, cellHeight int,
	sqStyle lipgloss.Style,
) string {
	if piece == "" {
		return sqStyle.Width(cellWidth).Height(cellHeight).Render(" ")
	}

	upper := strings.ToUpper(piece)
	var art []string

	switch {
	case cellWidth >= 8 && cellHeight >= 4:
		art = largePieces[upper]
	case cellWidth >= 6 && cellHeight >= 3:
		art = mediumPieces[upper]
	case cellWidth >= 5 && cellHeight >= 3:
		art = compactPieces[upper]
	default:
		glyph := unicodePieces[piece]
		return sqStyle.Width(cellWidth).Height(cellHeight).Render(glyph)
	}

	pieceColor := style.WhitePieceColor
	if piece == strings.ToLower(piece) {
		pieceColor = style.BlackPieceColor
	}

	lines := make([]string, cellHeight)
	topPad := (cellHeight - len(art)) / 2

	lineStyle := sqStyle.
		Width(cellWidth).
		UnsetHeight().
		Foreground(pieceColor).
		Bold(true)

	for r := 0; r < cellHeight; r++ {
		artIdx := r - topPad
		text := " "
		if artIdx >= 0 && artIdx < len(art) {
			text = art[artIdx]
		}
		lines[r] = lineStyle.Render(text)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

type boardLayout struct {
	cellWidth    int
	cellHeight   int
	sidebarWidth int
	stacked      bool // true on narrow terminals: put sidebar below board
}

func (m GameModel) calculateLayout() boardLayout {
	w := m.width
	h := m.height
	if w <= 0 {
		w = 120
	}
	if h <= 0 {
		h = 36
	}

	const (
		rankColWidth   = 2
		sidebarInnerW  = 24
		sidebarChromeW = 4 // border (2) + padding (2)
		verticalChrome = 8 // clocks + file labels + input + help line + margins
		minCellW       = 4
		maxCellW       = 10
		minCellH       = 2
		maxCellH       = 5
	)

	stacked := w < 76

	var sidebarWidth, availBoardW int
	if stacked {
		sidebarWidth = max(20, w-sidebarChromeW-4)
		availBoardW = w - rankColWidth - 4
	} else {
		sidebarWidth = sidebarInnerW
		availBoardW = w - rankColWidth - (sidebarWidth + sidebarChromeW) - 4
	}

	availBoardH := max(8, h-verticalChrome)
	if stacked {
		availBoardH = max(8, (h-verticalChrome)-12)
	}

	// Each square has 8 columns and 8 rows
	maxWPerCell := max(1, availBoardW/8)
	maxHPerCell := max(1, availBoardH/8)

	// Terminal cells are ~1:2 (width:height), so cellW ≈ 2 * cellH looks square.
	// Pick the largest balanced pair that fits in BOTH width and height budgets.
	cellH := clamp(min(maxHPerCell, maxWPerCell/2), minCellH, maxCellH)
	cellW := clamp(min(maxWPerCell, cellH*2), minCellW, maxCellW)

	// If height is generous (like h=51) but width is tight (like w=78),
	// keep cellH at least 3 as long as maxHPerCell >= 3, and let cellW use what's available.
	if maxHPerCell >= 3 && cellH < 3 {
		cellH = 3
		cellW = clamp(maxWPerCell, minCellW, maxCellW)
	}

	return boardLayout{
		cellWidth:    cellW,
		cellHeight:   cellH,
		sidebarWidth: sidebarWidth,
		stacked:      stacked,
	}
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
