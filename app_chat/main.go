package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/server"
	"github.com/raflyadek/chessTUI/internal/terminal"
)

var (
	host = "0.0.0.0"
	port = 3000
)

func main() {
	//ssh
	lobby := middleware.NewLobby()
	sshHost := server.NewSSHHost(host, port, lobby)
	defer sshHost.Stop()

	username := os.Getenv("USER")
	if username == "" {
		username = "Host"
	}

	root := terminal.InitRootModelWithLobby(lobby, username, port).WithStartHost(sshHost.Start)
	p := tea.NewProgram(root)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
		os.Exit(1)
	}
	// server.InitServer(host, port, middleware.NewLobby())
	// p := tea.NewProgram(terminal.InitRootModel())
	// if _, err := p.Run(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Oof: %v\n", err)
	// }
}
