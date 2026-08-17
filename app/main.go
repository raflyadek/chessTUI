package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	// "charm.land/bubbles/v2/textarea"
	// tea "charm.land/bubbletea/v2"
)

var moveCounter int = 0
var blackCastle bool = true
var whiteCastle bool = true
var enPassantMoveCounter int = 0
var enPassantRow int = 0
var enPassantCol int = 0
var isEnPassant bool = false

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

		(pawn)
		all the other i forgot to document it, but rn i got confused for the en passant
		how do we know if this move is available for en passant and which square/pawn?
	*/

	//init the board
	board := initBoard()

	notation := "    a    b    c    d    e    f    g    h    "
	//loop until checkmate or resign
	for {
		numberNotation := 8

		for i := 0; i < len(board); i++ {
			//print notatin at the top before the loop begin
			if i == 0 {
				fmt.Println(notation)
				fmt.Println()
			}
			fmt.Printf("%d ", numberNotation)
			for j := 0; j < len(board); j++ {
				if board[i][j] == "" {
					if (i+j)%2 == 0 {
						fmt.Print("  .  ")
					} else {
						fmt.Print("  ,  ")
					}
				} else {
					fmt.Print(board[i][j])
				}
				//add number notation at the end of loop
				if j == 7 {
					fmt.Printf(" %d", numberNotation)
					fmt.Println()
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

		fmt.Printf("Its %s move", player)
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
	board[7] = [8]string{"  R  ", "  N  ", "  B  ", "  Q  ", "  K  ", "  B  ", "  N  ", "  R  "}
	board[6] = [8]string{"  P  ", "  P  ", "  P  ", "  P  ", "  P  ", "  P  ", "  P  ", "  P  "}

	//black
	board[1] = [8]string{"  p  ", "  p  ", "  p  ", "  p  ", "  p  ", "  p  ", "  p  ", "  p  "}
	board[0] = [8]string{"  r  ", "  n  ", "  b  ", "  q  ", "  k  ", "  b  ", "  n  ", "  r  "}

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
	//fromCol take input index 0 for exmaple c and we substract it
	//with a then we get index 2 because decimal number of c is 97 and a is 95
	//then fromRow is we substract 8 with ascii decimal number of index 1 of input
	//for example is 2 and we subtract it again with 0 then we get 2 and we total it
	//8 - 2 = 6
	//so combining all of that c2 -> board[2][6]
	fromCol := int(from[0] - 'a')
	fromRow := 8 - int(from[1]-'0')
	toCol := int(to[0] - 'a')
	toRow := 8 - int(to[1]-'0')

	pieceLocation := strings.TrimSpace(board[fromRow][fromCol])

	board[toRow][toCol] = board[fromRow][fromCol]
	board[fromRow][fromCol] = ""

	// pieceDestination := board[toRow][toCol]
	//pawn promote

	//why board[toRow][toCol] works but if i put that into a variable
	//and used it, its not work??
	//promote pawn
	if pieceLocation == "P" || pieceLocation == "p" {
		if toRow == 7 || toRow == 0 {
			promote := promotePawn()
			board[toRow][toCol] = promote
		}
	}

	if (pieceLocation == "K" || pieceLocation == "k") && (toCol == fromCol+2 || toCol == fromCol-2) {
		//swap rook to beside king
		if toCol == fromCol+2 {
			board[toRow][toCol-1] = board[toRow][toCol+1]
			board[toRow][toCol+1] = ""
		}
		if toCol == fromCol-2 {
			board[toRow][toCol+1] = board[toRow][toCol-2]
			board[toRow][toCol-2] = ""
		}
	}
	if pieceLocation == "k" || pieceLocation == "K" || pieceLocation == "r" || pieceLocation == "R" {
		moveState(pieceLocation, fromRow, toRow, fromCol, toCol, board)
	}

	//en-passant pawn
	if pieceLocation == "p" || pieceLocation == "P" {
		moveBefore(fromRow, toCol, toRow, pieceLocation)
		if isEnPassant == true {
			board[enPassantRow][enPassantCol] = ""
			isEnPassant = false
		}
	}
	return board
}

/*
when input with the right notation, it still return false
*/
func legalMove(from, to string, board [8][8]string, moveCounter int) (bool, error) {
	notation := "abcdefgh"
	pieces := "prnbqk"

	//check if the from or to is > 2 character
	if len(from) > 2 || len(to) > 2 {
		return false, fmt.Errorf("invalid notation")
	} else if len(from) == 0 || len(to) == 0 {
		return false, fmt.Errorf("invalid notation")
	}

	//check if from/to is within the notation
	if fromCheck := strings.Contains(notation, string(from[0])); !fromCheck {
		return false, fmt.Errorf("the from notation is invalid")
	}

	if toCheck := strings.Contains(notation, string(to[0])); !toCheck {
		return false, fmt.Errorf("the to notation is invalid ")
	}

	//get input from and to index 1 to int and if its out of bound just throw false
	fromInt, err := strconv.Atoi(string(from[1:]))
	if err != nil {
		return false, err
	}
	if fromInt > 8 || fromInt < 1 {
		return false, fmt.Errorf("number notation is out of bounds")
	}

	toInt, err := strconv.Atoi(string(to[1:]))
	if err != nil {
		fmt.Printf("error parse to int %w", err)
		return false, err
	}
	if toInt > 8 || toInt < 1 {
		return false, fmt.Errorf("number notation is out of bound")
	}

	//check if the from is empty or not
	fromCol := int(from[0] - 'a')
	fromRow := 8 - int(from[1]-'0')
	toCol := int(to[0] - 'a')
	toRow := 8 - int(to[1]-'0')

	//get what piece it want to move
	pieceLocation := strings.TrimSpace(board[fromRow][fromCol])
	pieceDestination := strings.TrimSpace(board[toRow][toCol])
	// fmt.Printf("piece location: a%s  ", pieceLocation)
	// fmt.Printf("poece destination: b%s ", pieceDestination)

	if pieceLocation == "" {
		return false, fmt.Errorf("you cant move piece from empty square")
	}

	//check if the input is empty string
	if from == "" || to == "" {
		return false, fmt.Errorf("enter the right notation")
	}

	// if its white turn then its only can move the upper case pieces
	player := playerMove(moveCounter)
	if player == "White" {
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
	if strings.Contains(player, "White") {
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

	//piece cant move to row more than 7 or below than 0
	// if (pieceDestination)

	//pieces cant move past if there are piece in the middle destination

	//pieces rules?
	pErr := piecesRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
	if pErr != nil {
		return false, fmt.Errorf(pErr.Error())
	}
	// //check state for castle and en passant
	// whiteCastle, blackCastle, _, messageCastle := moveState(pieceLocation, fromRow, toRow, fromCol, toCol, board)
	// fmt.Println(messageCastle)
	// if blackCastle == false || whiteCastle == false {
	// 	if pieceLocation == "k" || pieceLocation == "K" && toCol != fromCol+1 {
	// 		return false, fmt.Errorf("you move king already, cant castle")
	// 	}
	// }
	// fmt.Println(messageCastle)
	// after all condition pased then return true
	return true, nil
}

/*
take turns white/black
*/
func playerMove(moveCounter int) string {
	if moveCounter%2 == 0 {
		return "White"
	} else {
		return "Black"
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
func piecesRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	// fmt.Println("fromrow: ", fromRow)
	// fmt.Println("fromCol", fromCol)
	// fmt.Println("from: ", from)
	// fmt.Println("to: ", to)
	// fmt.Println("toRow: ", toRow)
	// fmt.Println("toCol: ", toCol)
	// fmt.Println("piece location: ", pieceLocation)
	// fmt.Println("piece destination: ", pieceDestination)
	switch pieceLocation {
	case "p", "P":
		err := pawnRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
		if err != nil {
			return err
		}
	case "r", "R":
		err := rookRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
		if err != nil {
			return err
		}
	case "n", "N":
		err := knightRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow)
		if err != nil {
			return err
		}
	case "b", "B":
		err := bishopRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
		if err != nil {
			return err
		}
	case "q", "Q":
		err := queenRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
		if err != nil {
			return err
		}
	case "k", "K":
		err := kingRules(from, to, pieceLocation, pieceDestination, fromRow, fromCol, toCol, toRow, board)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("what piece is this? lol")
	}
	return nil
}

/*
pawn rules, its exactly what it sounds
*/
func pawnRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	differenceRowRaw := fromRow - toRow
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)
	//check is any en passant available
	//wtf is this messy logic lol
	if moveCounter == enPassantMoveCounter+1 {
		if fromCol != 0 && fromCol != 7 {
			besidePawn := strings.TrimSpace(board[fromRow][fromCol+1])
			besidePawn2 := strings.TrimSpace(board[fromRow][fromCol-1])

			if pieceLocation == "P" && toCol == enPassantCol && differenceRowAbs == 1 {
				if besidePawn == "p" || besidePawn2 == "p" {
					isEnPassant = true
					return nil
				}
			}
			if pieceLocation == "p" && toCol == enPassantCol && differenceRowAbs == 1 {
				if besidePawn == "P" || besidePawn2 == "P" {
					isEnPassant = true
					return nil
				}
			}
		}
		//for column a enpassant
		if fromCol == 0 {
			if pieceLocation == "P" && toCol == enPassantCol && differenceRowAbs == 1 {
				if strings.TrimSpace(board[fromRow][fromCol+1]) == "p" {
					isEnPassant = true
					return nil
				}
			}
			if pieceLocation == "p" && toCol == enPassantCol && differenceRowAbs == 1 {
				if strings.TrimSpace(board[fromRow][fromCol+1]) == "P" {
					isEnPassant = true
					return nil
				}
			}
		}
		//for column h enpassant
		if fromCol == 7 {
			if pieceLocation == "P" && toCol == enPassantCol && differenceRowAbs == 1 {
				if strings.TrimSpace(board[fromRow][fromCol-1]) == "p" {
					isEnPassant = true
					return nil
				}
			}
			if pieceLocation == "p" && toCol == enPassantCol && differenceRowAbs == 1 {
				if strings.TrimSpace(board[fromRow][fromCol-1]) == "P" {
					isEnPassant = true
					return nil
				}
			}
		}
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

	//if there are pieces in the destination location
	//and in the same notation, cant move forward
	if pieceDestination != "" && from[0] == to[0] {
		return fmt.Errorf("there is a piece, pawn cant move forward")
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

	//can only eat diagonal/column +1/-1 from its position
	if from[0] != to[0] {
		if pieceDestination != "" {
			if fromCol-toCol != 1 && fromCol-toCol != -1 {
				return fmt.Errorf("can only move 1 square to diagonal")
			}
		} else {
			return fmt.Errorf("move diagonal when there is a piece to eat")
		}
	}

	//en passant (this shit hard)
	/*
		i think its should be have their own function just like promotePawn()
		there is just too much to do, keeping a counter, etc
	*/
	// if pieceLocation == "p" && fromRow > 4 {

	return nil
}

// knight rules
func knightRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int) error {
	//i think its quite easy, its either move vertical or horizontal for + 2/-2
	//then just - 1/+1 after for either vertical/horizontal

	differenceColRaw := fromCol - toCol
	differenceRowRaw := fromRow - toRow
	differenceColAbs := max(differenceColRaw, -differenceColRaw)
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)
	if differenceColAbs > 2 || differenceRowAbs > 2 {
		return fmt.Errorf("knight can only jump 3 square")
	}
	//cant arrive at the same row/col
	if fromRow == toRow || fromCol == toCol {
		return fmt.Errorf("invalid knight move")
	}
	if differenceColAbs == 2 {
		if differenceRowAbs != 1 {
			return fmt.Errorf("thats invalid knight move")
		}
	}

	if differenceColAbs == 1 {
		if differenceRowAbs != 2 {
			return fmt.Errorf("invalid knight move man")
		}
	}
	//thats it no?
	return nil
}

// bishop rules
func bishopRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	//make sure bishop cant move vertical or horizontal
	if from[0] == to[0] || from[1] == to[1] {
		return fmt.Errorf("bishop must move diagonally")
	}

	//example:
	//if white player bishop white is in the d3, then it can move either to
	//c4 and so on or e4 or move backward to c2 or e2
	//we use a difference from where we standing
	differenceColRaw := fromCol - toCol
	differenceRowRaw := fromRow - toRow
	differenceColAbs := max(differenceColRaw, -differenceColRaw)
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)
	if differenceColAbs != differenceRowAbs {
		return fmt.Errorf("the move is ilegal")
	}

	//cant move if there is a piece blocking our way
	//if its the same color, then cant move past them
	//if its opposite color then we can eat, but cant move past them
	//we use difference too i think? and from there
	//get either 1 or -1
	rowDir := (toRow - fromRow) / int(math.Abs(float64(toRow)-float64(fromRow)))
	colDir := (toCol - fromCol) / int(math.Abs(float64(toCol)-float64(fromCol)))
	for i := 1; i < differenceColAbs; i++ {
		//check every step before the destination
		row := fromRow + i*rowDir
		col := fromCol + i*colDir
		if board[row][col] != "" {
			return fmt.Errorf("Illegal move there is a piece block your move")
		}
	}

	//i think thats it for bishop?
	return nil
}

// rook rules
func rookRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	// cant move diagonally, only vertical or horizontal
	//if its vertical then the column is the same from[0] == to[0]
	//if its horizontal then the row is the same from[1] == to[1]
	if from[1] != to[1] && from[0] != to[0] {
		return fmt.Errorf("Illegal move, rook move either horizontal or vertical")
	}

	//check every move bbefore the destination
	differenceColRaw := fromCol - toCol
	differenceRowRaw := fromRow - toRow
	differenceColAbs := max(differenceColRaw, -differenceColRaw)
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)
	//vertical
	if from[0] == to[0] {
		rowDir := (toRow - fromRow) / int(math.Abs(float64(toRow)-float64(fromRow)))
		for i := 1; i < differenceRowAbs; i++ {
			//check every step before destination
			row := fromRow + i*rowDir
			col := fromCol
			if board[row][col] != "" {
				return fmt.Errorf("Illegal move there is a piece blocking your way")
			}
		}
	}
	//horizontal
	if from[1] == to[1] {
		colDir := (toCol - fromCol) / int(math.Abs(float64(toCol)-float64(fromCol)))
		for i := 1; i < differenceColAbs; i++ {
			//check every step before destination
			row := fromRow
			col := fromCol + i*colDir

			if board[row][col] != "" {
				return fmt.Errorf("Illegal move there is a piece blocking your way")
			}
		}
	}
	//thats it no?
	return nil
}

// queen rules
func queenRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	differenceColRaw := fromCol - toCol
	differenceRowRaw := fromRow - toRow
	differenceColAbs := max(differenceColRaw, -differenceColRaw)
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)

	//queen move rules combine rook+bishop
	if from[1] != to[1] && from[0] != to[0] && differenceColAbs != differenceRowAbs {
		return fmt.Errorf("the move is ilegal")
	}

	//rook-like rules
	if from[0] == to[0] || from[1] == to[1] {
		//vertical
		if from[0] == to[0] {
			rowDir := (toRow - fromRow) / int(math.Abs(float64(toRow)-float64(fromRow)))
			for i := 1; i < differenceRowAbs; i++ {
				//check every step before destination
				row := fromRow + i*rowDir
				col := fromCol
				fmt.Printf("col: %d", col)
				if board[row][col] != "" {
					return fmt.Errorf("Illegal move there is a piece blocking your way")
				}
			}
		}
		//horizontal
		if from[1] == to[1] {
			colDir := (toCol - fromCol) / int(math.Abs(float64(toCol)-float64(fromCol)))
			for i := 1; i < differenceColAbs; i++ {
				//check every step before destination
				row := fromRow
				col := fromCol + i*colDir

				if board[row][col] != "" {
					return fmt.Errorf("Illegal move there is a piece blocking your way")
				}
			}
		}
	} else {
		//cant move if there is a piece blocking our way
		//if its the same color, then cant move past them
		//if its opposite color then we can eat, but cant move past them
		//we use difference too i think? and from there
		//get either 1 or -1
		//bishop-like rule
		rowDir := (toRow - fromRow) / int(math.Abs(float64(toRow)-float64(fromRow)))
		colDir := (toCol - fromCol) / int(math.Abs(float64(toCol)-float64(fromCol)))
		for i := 1; i < differenceColAbs; i++ {
			//check every step before the destination
			row := fromRow + i*rowDir
			col := fromCol + i*colDir
			if board[row][col] != "" {
				return fmt.Errorf("Illegal move there is a piece block your move")
			}
		}
	}

	//thats it no?
	return nil
}

// king rules
func kingRules(from, to, pieceLocation, pieceDestination string, fromRow, fromCol, toCol, toRow int, board [8][8]string) error {
	//its kinda easy except for the castle

	differenceColRaw := fromCol - toCol
	differenceRowRaw := fromRow - toRow
	differenceColAbs := max(differenceColRaw, -differenceColRaw)
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)

	//or king move +2 square col if its castle
	if differenceColAbs > 2 {
		return fmt.Errorf("king only move 1 square or 2 square for castle")
	}
	//king move only + 1 square for diagonal or vertical
	if differenceRowAbs >= 2 {
		return fmt.Errorf("king only move 1 square diagonal or vertical")
	}

	//castle
	if blackCastle == true || whiteCastle == true {
		if toCol == fromCol+2 {
			if board[fromRow][fromCol+1] != "" || board[fromRow][fromCol+2] != "" {
				return fmt.Errorf("cant castle, there is a piece block the way")
			}
			return nil
		}
		if toCol == fromCol-2 {
			if board[fromRow][fromCol-1] != "" || board[fromRow][fromCol-2] != "" || board[fromRow][fromCol-3] != "" {
				return fmt.Errorf("cant castle, there is a piece blocking the way")
			}
			return nil
		}
	}

	//just move anywhere but only +1 square
	//move +1 on vertical / horizontal / diaognal
	if blackCastle == false && pieceLocation == "k" {
		if to[1] != from[1]+1 && to[1] != from[1]-1 && to[0] != from[0]+1 && to[0] != from[0]-1 {
			return fmt.Errorf("king only move 1 square")
		}
	}

	if whiteCastle == false && pieceLocation == "K" {
		if to[1] != from[1]+1 && to[1] != from[1]-1 && to[0] != from[0]+1 && to[0] != from[0]-1 {
			return fmt.Errorf("king only move 1 square")
		}
	}
	//checkmate/check or open check from another piece
	//or create another function to check every move if that move
	//threaten the king or if that move open check the king and return
	//string maybe -> "Check/Open Check"

	// i think thats it?
	return nil
}

func promotePawn() string {
	//promote
	pieces := "rnbq"

	var piecesPromote string

	if playerMove(moveCounter) == "White" {
		piecesPromote = "(R/N/B/Q)"
	} else {
		piecesPromote = "(r/n/b/q)"
	}

	var promote string

	for {
		fmt.Printf("select pieces you want to promote %s: ", piecesPromote)
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		promote = reader.Text()

		if playerMove(moveCounter) == "White" {
			if !strings.Contains(strings.ToUpper(pieces), promote) {
				fmt.Println("select the right pieces")
			} else {
				return promote
			}
		} else {
			if !strings.Contains(strings.ToLower(pieces), promote) {
				fmt.Println("select the right piece")
			} else {
				return promote
			}
		}
	}
}

/*
	TODO:

checkmate/check or open check from another piece
or create another function to check every move if that move
threaten the king or if that move open check the king and return
string maybe -> "Check/Open Check"
*/
func checkMove(pieceLocation string) (string, error) {
	//check if the next move from previous move is threaten the king
	//if bishop then check possible move for bishop if that can move to king
	//if yes then it check the opposite king or if we want to move
	//
	//or we can check the entire board and determine which piece can't move
	//because they block the king from check

	return "", nil
}

/*
this function to track the move before, it purpose is to check
if the king still can castle or not, and the if the pawn can perform en passant
and we can put the function just before apply move?

MAYBE we can just create a fen notation tho? i think its easier to save the state then
we just need to transfer the whole board into fen and save it to map[moveCounter] = fen

//TODO: Enpassant

	if piecelOCATION == p/P and it moved + 2 from starting point
	then check if there is a +1/-1 column there is opposite pawn or no
	if yes then enPassant = true and it saved col/row pawn that are eligible
	white pawn must be at the exactly row 5/index 3
*/
func moveBefore(fromRow, toCol, toRow int, pieceLocation string) (int, int, int) {
	differenceRowRaw := fromRow - toRow
	differenceRowAbs := max(differenceRowRaw, -differenceRowRaw)

	//the && binds tighter than ||, so used parentheses in the || not &&
	if (pieceLocation == "P" || pieceLocation == "p") && differenceRowAbs == 2 {
		enPassantCol = toCol
		enPassantRow = toRow
		enPassantMoveCounter = moveCounter
	}
	//if there is a pawn move 2 square save the col and row and moveCounter
	return enPassantCol, enPassantRow, enPassantMoveCounter
}
func moveState(pieceLocation string, fromRow, toRow, fromCol, toCol int, board [8][8]string) {
	//why it is update when the move is invalid like e8 to b8 but it update the state?
	if playerMove(moveCounter) == "Black" && (board[0][4] != "k" || board[0][0] != "r" || board[0][7] != "r") {
		blackCastle = false
	}
	if playerMove(moveCounter) == "White" && (board[7][4] != "K" || board[7][0] != "R" || board[7][7] != "R") {
		whiteCastle = false
	}
}
