package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	//task 4. put the pieces in the right square (done)
	//task 5. make the pieces can move, ignoring all the rules first
	//task 6. make each player move white/black
	//task 7. make pieces move rule

	/*
		(task 1) (done)
		in outer loop we do nothing, just jump straight to inner loop
		which is the column, so after each outer loop done (for the row we
		jump to the column and fill it with either i or j)

		because outer loop done after first loop, meanwhile inner loop
		done after it reach the condition (j < len(board))
	*/

	/*
		(task 2) (done)
		because  j represent the column, and i is row, then after
		the last index of j it is the start of the new row, just add
		"\n" or escape sequence so it add new line, and to print it without
		"[ ]", simply just print board[i][j] at the end of inner loop finnished
	*/

	/*
		(task 3) (done)
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
		(task 4) (done)
		pieces representation is kinda tricky, i think we need to hardcode it first in the right notation,
		and after it whenever its move render it again, not render it in the game first start

		not that tricky tho, just create a function (initBoard) with a board is prefilled with pieces, and then
		loop the array again if the element is equal to "" <- because we use string, then its cannot be null
		then we check if element equal to empty string / "", then if its empty then, print the board representation
		(,) as black and (.) as white
	*/

	/*
		(task 5) (partial? will definitely add more logic to this function)
		create a function to take an input move from user (piecesMove function) and pass it as a return
		then create a function to apply that move (applyMove) that takes an input from piecesMove function return

		maybe create a function to check if the move is legal/not? like func legalMove() [8][8]string
		and if we take that route with the function returning the board, then we can show whats is legal,
		like if i put e4 and enter then it will show the dot to mark the legal move of the pieces,
			(done still simple only if the notation is out of bounds)
		OR
		we can legalMove() bool <- then we can just check if the move is legal or not, so when the user want to move
		e4 to e16, its immediately return false and show the message "the move is not legal"
	*/

	/*
		(task 6) (done)
		each player move, start with white and then black and repeat until it checkmate.

		we ccan create a function (playerMove) that receive a parameter of counter
		that return the the string if its a white or black
		by using the counter if the counter % 2 == 0 then its white else black
	*/

	/*
		(task 7) (progress)
		piece move rule

		its either create one function for all the pieces
		or
		create a separate function for every piece
		-> its better to create a separate function for every piece because
		if there are adjustment its not clunky on the function, we can achieve
		separate of concern, rather than one function, the code will be messy

	*/

	//init the board
	board := initBoard()

	moveCounter := 0
	notation := "  abcdefgh"
	//loop until checkmate or resign
	for {
		numberNotation := 8

		for i := 0; i < len(board); i++ {
			//print notatin at the top before the loop begin
			if i == 0 {
				fmt.Println(notation)
			}
			fmt.Printf("%d ", numberNotation)
			for j := 0; j < len(board); j++ {
				if board[i][j] == "" {
					if (i+j)%2 == 0 {
						fmt.Print(".")
					} else {
						fmt.Print(",")
					}
				} else {
					fmt.Print(board[i][j])
				}
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
				fmt.Println()
			}
		}
		player := playerMove(moveCounter)

		/*
			blocker: if we put piecesMove here, the variable cannot be used to generate a move because
			board representation is above this, how do i use this variable? (done)
			just create another function that return [8][8]string and use that as a new board
		*/

		fmt.Print(player)
		from, to := piecesMove()
		flag, err := legalMove(from, to, board, moveCounter)

		if flag == true {

			board = applyMove(board, from, to)
			// add counter if only the move is legal counter for move
			moveCounter++
		} else {
			fmt.Printf("error: %s\n", err)
			fmt.Println()
		}
	}
}

/*
initial board when play
*/
func initBoard() [8][8]string {
	board := [8][8]string{}

	//white
	board[7] = [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	board[6] = [8]string{"P", "P", "P", "P", "P", "P", "P", "P"}

	//black
	board[1] = [8]string{"p", "p", "p", "p", "p", "p", "p", "p"}
	board[0] = [8]string{"r", "n", "b", "q", "k", "b", "n", "r"}

	return board
}

/*
ignore all legal just pieces can move first
already have a function to make sure its legal (check only the notation out of bounds or not)
*/
func piecesMove() (from, to string) {
	fmt.Println("\n(Insert the notation)")
	fmt.Print("What piece you want to move? ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	from = reader.Text()
	fmt.Print("To what square? ")
	reader.Scan()
	to = reader.Text()
	fmt.Println()

	return from, to
}

/*
apply move from input piecesMove()
*/
func applyMove(board [8][8]string, from, to string) [8][8]string {
	//from e2 to e3'
	fromCol := int(from[0] - 'a')
	fromRow := 8 - int(from[1]-'0')
	toCol := int(to[0] - 'a')
	toRow := 8 - int(to[1]-'0')

	// fmt.Println(fromCol, fromRow, toCol, toRow)
	//copy piece to (to) and delete it from (from)
	board[toRow][toCol] = board[fromRow][fromCol]
	board[fromRow][fromCol] = ""

	return board
}

/*
	when input with the right notation, it still return false

TODO: we can change the return to err and change the return in each logic to
fmt.ErrorF and send the message to err, so we can show the err message
to the output instead of just "illegal move"
*/
func legalMove(from, to string, board [8][8]string, moveCounter int) (bool, error) {
	notation := "abcdefgh"
	pieces := "prnbqk"
	//check if from/to is within the notation
	if fromCheck := strings.Contains(notation, string(from[0])); !fromCheck {
		return false, fmt.Errorf("the from notation is invalid")
	}

	if toCheck := strings.Contains(notation, string(to[0])); !toCheck {
		return false, fmt.Errorf("the to notation is invalid ")
	}

	//check if the from is empty or not
	fromCol := int(from[0] - 'a')
	fromRow := 8 - int(from[1]-'0')
	toCol := int(to[0] - 'a')
	toRow := 8 - int(to[1]-'0')

	//get what piece it want to move
	pieceLocation := board[fromRow][fromCol]
	pieceDestination := board[toRow][toCol]
	// fmt.Println("[oece location: ", pieceLocation)
	// fmt.Println("[oece destination: ", pieceDestination)

	if pieceLocation == "" {
		return false, fmt.Errorf("you cant move piece from empty square")
	}

	//check if the input is empty string
	if from == "" || to == "" {
		return false, fmt.Errorf("enter the right notation")
	}

	//get input from and to index 1 to int and if its out of bound just throw false
	fromInt, err := strconv.Atoi(string(from[1:]))
	if err != nil {
		return false, err
	}
	if fromInt > 8 {
		return false, fmt.Errorf("number notation is out of bounds")
	}

	toInt, err := strconv.Atoi(string(to[1:]))
	if err != nil {
		fmt.Printf("error parse to int %w", err)
		return false, err
	}
	if toInt > 8 {
		return false, fmt.Errorf("number notation is out of bound")
	}

	// if its white turn then its only can move the upper case pieces
	player := playerMove(moveCounter)
	if strings.Contains(player, "white") {
		if !strings.Contains(strings.ToUpper(pieces), pieceLocation) {
			return false, fmt.Errorf("cant move black piece when its white turn")
		}
	} else {
		if !strings.Contains(strings.ToLower(pieces), pieceLocation) {
			return false, fmt.Errorf("cant move white piece when its black turn")
		}
	}

	/*
		cant eat piece with the same color

		the guard pieceDestination != "" is because every string contains empty string,
		so when there is no guard, the actual logic condition always return true
	*/
	if strings.Contains(player, "white") {
		if pieceDestination != "" {
			if strings.Contains(strings.ToUpper(pieces), pieceLocation) == strings.Contains(strings.ToUpper(pieces), pieceDestination) {
				return false, fmt.Errorf("white piece cant eat white piece")
			}
		}
	} else {
		if pieceDestination != "" {
			if strings.Contains(strings.ToLower(pieces), pieceLocation) == strings.Contains(strings.ToLower(pieces), pieceDestination) {
				return false, fmt.Errorf("black piece cant eat black piece")
			}
		}
	}

	//pieces rules?
	pErr := piecesRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow)
	if pErr != nil {
		return false, fmt.Errorf(pErr.Error())
	}

	// after all condition pased then return true
	return true, nil
}

/*
take turns white/black
*/
func playerMove(moveCounter int) string {
	if moveCounter%2 == 0 {
		return "Its white turn"
	} else {
		return "Its black turn"
	}
}

/*
rules for every pieces

is it better to make a different function for different piece
or
is it better to make a single function that validate every piece

ok, make a different function for different piece and later maybe we create one function
to validate the from pieces, and if p then it goes to pawnRules() function, etc.
*/
func piecesRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol,toCol, toRow int) error {
	fmt.Println("fromrow: ", fromRow)
	fmt.Println("fromCol", fromCol)
	fmt.Println("from: ", from)
	fmt.Println("to: ", to)
	fmt.Println("piece location: ", pieceLocation)
	fmt.Println("piece destination: ", pieceDestination)
	switch pieceLocation {
	case "p", "P":
		err := pawnRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow)
		if err != nil {
			return err
		}
	case "r", "R":
		fmt.Println("this is rook")
	case "n", "N":
		fmt.Println("this is knight")
	case "b", "B":
		fmt.Println("this is bishop")
	case "q", "Q":
		fmt.Println("this is queen")
	case "k", "K":
		fmt.Println("this is king")
		// default:
		// 	return fmt.Errorf("what piece is this? lol")
	}

	return nil
}

func pawnRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int) error {
	//if there are pieces in the destination location 
	//and in the same notation, cant move forward
	if pieceDestination != "" && from[0] == to[0] {
		return fmt.Errorf("there is a piece, pawn cant move forward")
	}

	//can only move 1 or 2 square when never move before
	if pieceLocation == "P" && fromRow == 6 {
		if to[1] != from[1]+1 && to[1] != from[1]+2 {
			return fmt.Errorf("pawn white can only move 1 or 2 square in its starting position")
		}
	}

	if pieceLocation == "p" && fromRow == 1 {
		if to[1] != from[1]-1 && to[1] != from[1]-2 {
			return fmt.Errorf("pawn can only move 1 or 2 square in its starting position")
		}
	}

	//can only move 1 square if already move before
	if pieceLocation == "P" && fromRow != 6 {
		if to[1] != from[1]+1 {
			return fmt.Errorf("pawn only move 1 square if already move before")
		}
	}

	if pieceLocation == "p" && fromRow != 1 {
		if to[1] != from[1]-1 {
			return fmt.Errorf("pawn only move 1 square if already move before")
		}
	}
	fmt.Println("lewat")
	//can only eat diagonal/column +1/-1 from its position
	if from[0] != to[0] { 
		if pieceDestination != "" {
			if fromCol - toCol != 1 && fromCol - toCol != -1 {
				return fmt.Errorf("can only move 1 square to diagonal")
			}
		} else {
			return fmt.Errorf("move diagonal when there is a piece to eat")
		}
	}

	//en passant (this shit hard)

	return nil
}
