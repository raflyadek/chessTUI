package main

import (
	"fmt"
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
	//task 3. change i and j to white and black like a board representation
	//the board representation is complete with the notation aswell (a - h) & (1-8) (done)
	//task 4. put the pieces in the right square
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
		for board representation with black and white and pure using golang
		i dont think it achieveable? i think it can be achieveable when we already
		implement a bubbletea? we can ahieve it with (.) for white and (,) for black
		thats just temporary

		board representation is done using (.) as a white and (,) as a black, and for notation
		we use hardcode string and println it before the loop begin and after the loop end, for number notation
		we define a int with 8 and decrement it after at the end of column and we print it before the board print begin
		and after it finish and add a  (space) for cleanliness
	*/

	/*
		(task 4)
		pieces representation is kinda tricky, i think we need to hardcode it first in the right notation,
		and after it whenever its move render it again, not render it in the game first start
	*/

	notation := "  abcdefgh"
	numberNotation := 8
	for i := 0; i < len(board); i++ {
		//print notatin at the top before the loop begin
		if i == 0 {
			fmt.Println(notation)
		}
		fmt.Printf("%d ", numberNotation)
		for j := 0; j < len(board); j++ {
			if j%2 == 0 {
				board[i][j] = "."
			} else {
				board[i][j] = ","
			}
			fmt.Print(board[i][j])
			//add number notation at the end of loop
			if j == 7 {
				fmt.Printf(" %d", numberNotation)
				numberNotation--
			}
		}
		//add new line after every row
		fmt.Println()
		//print notatin at the top after the loop for row end
		if i == 7 {
			fmt.Println(notation)
		}
	}
}
