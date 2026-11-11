package main

import engine "github.com/raflyadek/chessTUI/internal/engine"

//TODO: ui with bubbletea
//TODO: multiplayer/ssh server

func main() {
	//init game state
	state := engine.NewGameState()
	//init the board
	board := engine.InitBoard()
	engine.PlayGame(state, board)
}
