package middleware

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	tea "charm.land/bubbletea/v2"
)

type Side string

const (
	SideWhite  Side = "white"
	SideBlack  Side = "black"
	SideRandom Side = "random"
)

type TimeFormat struct {
	Label     string
	Initial   time.Duration
	Increment time.Duration
}

var TimePresets = []TimeFormat{
	{Label: "1+0 Bullet", Initial: 1 * time.Minute, Increment: 0},
	{Label: "3+2 Blitz", Initial: 3 * time.Minute, Increment: 2 * time.Second},
	{Label: "5+0 Blitz", Initial: 5 * time.Minute, Increment: 0},
	{Label: "10+0 Rapid", Initial: 10 * time.Minute, Increment: 0},
}

// Messages pushed into each player's Bubble Tea event channel
type MatchStartedMsg struct {
	RoomID     string
	Assigned   Side
	TimeFormat TimeFormat
	Board      [8][8]string
}

type MoveUpdatedMsg struct {
	Board   [8][8]string
	Move    string
	Next    Side
	History []string
}

type OpponentDisconnectedMsg struct{}

type Player struct {
	Name   string
	Side   Side
	Events chan tea.Msg
}

type Room struct {
	mu         sync.Mutex
	ID         string
	TimeFormat TimeFormat
	HostSide   Side
	White      *Player
	Black      *Player
	Started    bool
	Turn       Side
	Board      [8][8]string
	History    []string
	// ApplyMove wraps your engine's `state.PlayGameTUI(board, move)`
	ApplyMove func(board [8][8]string, move string) ([8][8]string, error)
}

type Lobby struct {
	mu    sync.Mutex
	rooms map[string]*Room
}

func NewLobby() *Lobby {
	return &Lobby{
		rooms: make(map[string]*Room),
	}
}

func (l *Lobby) CreateRoom(
	hostName string,
	preferredSide Side,
	tf TimeFormat,
	initialBoard [8][8]string,
	newApplyFn func([8][8]string, string) ([8][8]string, error),
) (*Room, *Player) {
	l.mu.Lock()
	defer l.mu.Unlock()

	code := generateRoomCode()
	side := resolveSide(preferredSide)

	host := &Player{
		Name:   hostName,
		Side:   side,
		Events: make(chan tea.Msg, 16),
	}

	room := &Room{
		ID:         code,
		TimeFormat: tf,
		HostSide:   side,
		Board:      initialBoard,
		Turn:       SideWhite,
		ApplyMove:  newApplyFn,
	}

	if side == SideWhite {
		room.White = host
	} else {
		room.Black = host
	}

	l.rooms[code] = room
	return room, host
}

func (l *Lobby) JoinRoom(code, guestName string) (*Room, *Player, error) {
	l.mu.Lock()
	room, ok := l.rooms[strings.ToUpper(strings.TrimSpace(code))]
	l.mu.Unlock()

	if !ok {
		return nil, nil, fmt.Errorf("room %q not found", code)
	}

	room.mu.Lock()
	defer room.mu.Unlock()

	if room.Started || (room.White != nil && room.Black != nil) {
		return nil, nil, errors.New("room is already full")
	}

	guestSide := SideBlack
	if room.White == nil {
		guestSide = SideWhite
	}

	guest := &Player{
		Name:   guestName,
		Side:   guestSide,
		Events: make(chan tea.Msg, 16),
	}

	if guestSide == SideWhite {
		room.White = guest
	} else {
		room.Black = guest
	}
	room.Started = true

	// Notify both players to enter GameView
	room.White.Events <- MatchStartedMsg{
		RoomID:     room.ID,
		Assigned:   SideWhite,
		TimeFormat: room.TimeFormat,
		Board:      room.Board,
	}
	room.Black.Events <- MatchStartedMsg{
		RoomID:     room.ID,
		Assigned:   SideBlack,
		TimeFormat: room.TimeFormat,
		Board:      room.Board,
	}

	return room, guest, nil
}

func (r *Room) Leave(p *Player) {
	if r == nil || p == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	var other *Player
	if r.White == p {
		r.White = nil
		other = r.Black
	} else if r.Black == p {
		r.Black = nil
		other = r.White
	}

	if other != nil {
		select {
		case other.Events <- OpponentDisconnectedMsg{}:
		default:
		}
	}
}

func (r *Room) SubmitMove(playerSide Side, move string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.Started {
		return errors.New("waiting for opponent")
	}
	if r.Turn != playerSide {
		return errors.New("not your turn")
	}

	nextBoard, err := r.ApplyMove(r.Board, move)
	if err != nil {
		return err
	}

	r.Board = nextBoard
	r.History = append(r.History, move)
	if r.Turn == SideWhite {
		r.Turn = SideBlack
	} else {
		r.Turn = SideWhite
	}

	msg := MoveUpdatedMsg{
		Board:   r.Board,
		Move:    move,
		Next:    r.Turn,
		History: append([]string(nil), r.History...),
	}

	if r.White != nil {
		r.White.Events <- msg
	}
	if r.Black != nil {
		r.Black.Events <- msg
	}
	return nil
}

func (l *Lobby) RemoveRoom(code string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.rooms, code)
}

func resolveSide(s Side) Side {
	if s == SideWhite || s == SideBlack {
		return s
	}
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err == nil && n.Int64() == 1 {
		return SideBlack
	}
	return SideWhite
}

func generateRoomCode() string {
	const alphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ" // avoids 0/O and 1/I confusion
	b := make([]byte, 5)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		b[i] = alphabet[n.Int64()]
	}
	return string(b)
}

func (l *Lobby) JoinAvailableRoom(guestName string) (*Room, *Player, Side, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, room := range l.rooms {
		room.mu.Lock()
		var assigned Side
		if room.White == nil && room.Black != nil {
			assigned = SideWhite
		} else if room.Black == nil && room.White != nil {
			assigned = SideBlack
		} else {
			room.mu.Unlock()
			continue
		}

		guest := &Player{
			Name:   guestName,
			Side:   assigned,
			Events: make(chan tea.Msg, 16),
		}

		if assigned == SideWhite {
			room.White = guest
		} else {
			room.Black = guest
		}
		room.Started = true
		// Notify the host player waiting in SettingModel
		var host *Player
		if assigned == SideWhite && room.Black != nil {
			host = room.Black
		} else if assigned == SideBlack && room.White != nil {
			host = room.White
		}
		board := room.Board
		tf := room.TimeFormat
		room.mu.Unlock()

		if host != nil {
			select {
			case host.Events <- MatchStartedMsg{
				RoomID: room.ID,
				// Opponent:   guest,
				Assigned:   room.HostSide,
				TimeFormat: tf,
				Board:      board,
			}:
			default:
			}
		}

		return room, guest, assigned, true
	}

	return nil, nil, "", false
}

func (l *Lobby) GetRoom(code string) (*Room, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	room, ok := l.rooms[strings.ToUpper(strings.TrimSpace(code))]
	return room, ok
}
