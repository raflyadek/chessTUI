package main

import (
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/server"
)

var (
	host = "0.0.0.0"
	port = 3000
)

func main() {
	//ssh
	server.InitServer(host, port, middleware.NewLobby())
	// p := tea.NewProgram(terminal.InitRootModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	// }
}
