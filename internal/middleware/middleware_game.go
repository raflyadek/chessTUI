package middleware

import (
	"charm.land/ssh"
	"charm.land/wish/v2"
)

// TODO:
// create a room, player and gameMiddleware when user choose
// "play with a friend" menu and just use the view on game menu later

type Player struct {
	Name  string
	Piece chan string
}

type Room struct {
	players map[ssh.Session]Player
}

var room Room

func GameMiddleware() wish.Middleware {
	if len(room.players) == 0 {
		room = Room{
			players: make(map[ssh.Session]Player),
		}
	}
	return nil
}
