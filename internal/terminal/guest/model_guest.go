package guest

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
)

// BoardOnlyModel serves only the game board for a LAN guest session.
type BoardOnlyModel struct {
	game game.GameModel
}

func InitBoardOnlyModel(
	room *middleware.Room,
	player *middleware.Player,
	assigned middleware.Side,
	tf middleware.TimeFormat,
	board [8][8]string,
) BoardOnlyModel {
	return BoardOnlyModel{
		game: game.InitOnlineGameModel(room, player, assigned, tf, board),
	}
}

func (m BoardOnlyModel) Init() tea.Cmd {
	return m.game.Init()
}
