package main

import internal "github.com/raflyadek/chessTUI/internal/engine"

//TODO: ui with bubbletea
//TODO: multiplayer/ssh server

func main() {
	//init game state
	state := internal.NewGameState()
	//init the board
	board := internal.InitBoard()
	internal.PlayGame(state, board)
}
