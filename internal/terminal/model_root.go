package terminal

import (
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/engine"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/terminal/game"
	"github.com/raflyadek/chessTUI/internal/terminal/menu"
	"github.com/raflyadek/chessTUI/internal/terminal/setting"
)

type stateView int

const (
	menuView stateView = iota
	gameView
	settingView
)

// root model
type RootModel struct {
	width   int
	height  int
	state   stateView
	game    game.GameModel
	menu    menu.MenuModel
	setting setting.SettingModel
	lobby   *middleware.Lobby
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func InitRootModel() RootModel {
	return RootModel{
		menu: menu.InitMenuModel(),
		game: game.InitGameModel(),
	}
}

func InitRootModelWithLobby(
	lobby *middleware.Lobby,
	playerName string,
	sshPort int,
) RootModel {
	if playerName == "" {
		playerName = "Player"
	}

	g := game.InitGameModel()

	return RootModel{
		state:   menuView,
		lobby:   lobby,
		game:    g,
		menu:    menu.InitMenuModel(),
		setting: setting.InitSettingModel(lobby, playerName, sshPort, engine.InitBoard(), engine.NewGameState().PlayGameTUI),
	}
}
