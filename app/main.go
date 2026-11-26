package main

import engine "github.com/raflyadek/chessTUI/internal/engine"

func main() {
	//init game state
	state := engine.NewGameState()
	//init the board
	board := engine.InitBoard()
	engine.PlayGame(state, board)
}
