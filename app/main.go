package main

import (
	"fmt"
	"time"
	// "charm.land/bubbles/v2/textarea"
	// tea "charm.land/bubbletea/v2"
)

func main() {
	//in order to create chess i think we can use 2d array?
	//first array is for rows
	//second array is for column

	//task 1. print i j i j i j i j for every row (done)
	//task 2. print the array in a chess board square form, 
	//after each row then println/escape sequence ("\n") (done)
	//task 3. 
	board := [8][8]string{}

	/*
		(task 1)
		in outer loop we do nothing, just jump straight to inner loop
		which is the column, so after each outer loop done (for the row we 
		jump to the column and fill it with either i or j)

		because outer loop done after first loop, meanwhile inner loop  
		done after it reach the condition (j < len(board))
	*/
	/*
		(task 2)
		because  j represent the column, and i is row, then after 
		the last index of j it is the start of the new row, just add 
		"\n" or escape sequence so it add new line, and to print it without
		"[ ]", simply just print board[i][j] at the end of inner loop finnished 
	*/

	/*
		(task 3)
		.		
	*/
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if j % 2 == 0 {
				board[i][j] = "i"
			} else {
				board[i][j] = "j"
			}
			if j == 7 {
				board[i][j] += "\n"
			}
			fmt.Print(board[i][j])
			// time.Sleep(2 * time.Second)
		}
	}
}
