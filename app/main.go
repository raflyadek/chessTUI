package main

import engine "github.com/raflyadek/chessTUI/internal/engine"

func main() {
	//init game state
	state := engine.NewGameState()
	//init the board
	board := state.InitBoard()
	engine.PlayGame(state, board)
}
