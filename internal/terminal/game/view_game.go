package game

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

// delegate view
var ()

func (m GameModel) View() string {
	board := m.boardRender()
	blackTimer := m.blackTimerRender()
	whiteTimer := m.whiteTimerRender()
	moveInput := m.InputRender()
	// player := m.state.PlayerMove()

	playerState := m.player
	player := ""
	if playerState == whitePlayer {
		player = "whitE"
	} else {
		player = "Black"
	}
	s := m.err

	from := "from " + m.from + "|"
	to := "to " + m.to + "|"
	input := "input " + m.input
	return lipgloss.JoinVertical(lipgloss.Top, s, blackTimer, board, whiteTimer, input, from, to, player, moveInput)
}

// render board
func (m GameModel) boardRender() string {
	var s strings.Builder
	notation := "    a    b    c    d    e    f    g    h    "
	numberNotation := 8
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
	for i := 0; i < len(m.board); i++ {
		//print notatin at the top before the loop begin
		if i == 0 {
			s.WriteString(notation)
			s.WriteString("\n\n")
		}
		s.WriteString(fmt.Sprintf("%d ", numberNotation))
		for j := 0; j < len(m.board[i]); j++ {
			if m.board[i][j] == "" {
				if (i+j)%2 == 0 {
					s.WriteString("  .  ")
				} else {
					s.WriteString("  ,  ")
				}
			} else {
				s.WriteString(fmt.Sprintf("  %s  ", pieces[m.board[i][j]]))
			}
			//add number notation at the end of loop
			if j == 7 {
				s.WriteString(fmt.Sprintf(" %d\n", numberNotation))
				numberNotation--
			}
		}
		//add new line after every row
		s.WriteString("\n")
		//print notatin at the top after the loop for row end

		if i == 7 {
			s.WriteString(notation)
			s.WriteString("\n")
		}
	}
	return s.String()
}

// render clock
func (m GameModel) blackTimerRender() string {
	//timer active is from m.state.moveCounter
	//if %2==0 then white timer is on

	s := m.blackTimer.View()
	return s
}

func (m GameModel) whiteTimerRender() string {
	s := m.whiteTimer.View()
	return s
}

// render move input
func (m GameModel) InputRender() string {
	s := m.moveInput.View()
	return s
}

// render history
func (m GameModel) historyRender() string {

	return ""
}
