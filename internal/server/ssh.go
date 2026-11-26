package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/ssh"
	"charm.land/wish/v2"
	wishbm "charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/terminal"
)

func InitServer(host string, port int, lobby *middleware.Lobby) {
	server, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		//accept any public key.
		ssh.PublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),
		ssh.PasswordAuth(func(ssh.Context, string) bool { return false }),
		wish.WithMiddleware(
			//middleware for chess tui
			wishbm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				return terminal.InitRootModelWithLobby(lobby, s.User(), port), nil
			}),
			logging.Middleware(),
		),
		wish.WithIdleTimeout(1*time.Hour),
	)
	if err != nil {
		log.Fatalf("Couldn't create the server: %s", err)
	}
	startServer(server, err, host, port)
}

func startServer(server *ssh.Server, err error, host string, port int) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = server.ListenAndServe(); err != nil {
			log.Fatalln(err)
			done <- nil
		}
	}()
	<-done
	log.Printf("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Fatalln("Could not stop server", "error", err)
	}
}
