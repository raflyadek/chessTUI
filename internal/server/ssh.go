package server

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/ssh"
	"charm.land/wish/v2"
	wishbm "charm.land/wish/v2/bubbletea"
	"charm.land/wish/v2/logging"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/terminal"
	friend "github.com/raflyadek/chessTUI/internal/terminal/guest"
)

//async start

type SSHHost struct {
	mu      sync.Mutex
	srv     *ssh.Server
	running bool
	host    string
	port    int
	lobby   *middleware.Lobby
}

func NewSSHHost(host string, port int, lobby *middleware.Lobby) *SSHHost {
	return &SSHHost{
		host:  host,
		port:  port,
		lobby: lobby,
	}
}

// Start launches the Wish SSH server in the background if it isn't already running.
func (h *SSHHost) Start() error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.running {
		return nil
	}

	srv, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", h.host, h.port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		ssh.PublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),
		ssh.PasswordAuth(func(ssh.Context, string) bool { return false }),
		wish.WithMiddleware(
			wishbm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				guestName := s.User()
				if guestName == "" {
					guestName = "Friend"
				}
				room, guest, assigned, ok := h.lobby.JoinAvailableRoom(guestName)
				if !ok {

					// Incoming SSH friends open straight into the Join/Lobby flow or root menu
					return terminal.InitRootModelWithLobby(h.lobby, s.User(), h.port), nil
				}
				return friend.InitBoardOnlyModel(room, guest, assigned, room.TimeFormat, room.Board), nil
			}),
			logging.Middleware(),
		),
		wish.WithIdleTimeout(1*time.Hour),
	)
	if err != nil {
		return err
	}

	h.srv = srv
	h.running = true

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			h.mu.Lock()
			h.running = false
			h.mu.Unlock()
		}
	}()

	return nil
}

func (h *SSHHost) Stop() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.running || h.srv == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = h.srv.Shutdown(ctx)
	h.running = false
}

// func InitServer(host string, port int, lobby *middleware.Lobby) {
// 	server, err := wish.NewServer(
// 		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
// 		wish.WithHostKeyPath(".ssh/id_ed25519"),
// 		//accept any public key.
// 		ssh.PublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),
// 		ssh.PasswordAuth(func(ssh.Context, string) bool { return false }),
// 		wish.WithMiddleware(
// 			//middleware for chess tui
// 			wishbm.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
// 				return terminal.InitRootModelWithLobby(lobby, s.User(), port), nil
// 			}),
// 			logging.Middleware(),
// 		),
// 		wish.WithIdleTimeout(1*time.Hour),
// 	)
// 	if err != nil {
// 		log.Fatalf("Couldn't create the server: %s", err)
// 	}
// 	startServer(server, err, host, port)
// }
//
// func startServer(server *ssh.Server, err error, host string, port int) {
// 	done := make(chan os.Signal, 1)
// 	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
// 	log.Printf("Starting SSH server on %s:%d", host, port)
// 	go func() {
// 		if err = server.ListenAndServe(); err != nil {
// 			log.Fatalln(err)
// 			done <- nil
// 		}
// 	}()
// 	<-done
// 	log.Printf("Stopping SSH server")
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer func() { cancel() }()
// 	if err := server.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
// 		log.Fatalln("Could not stop server", "error", err)
// 	}
// }
