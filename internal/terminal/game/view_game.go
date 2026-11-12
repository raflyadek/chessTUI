package game

import (
	"fmt"
	"strings"
)

//delegate view

func (m GameModel) View() string {
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
