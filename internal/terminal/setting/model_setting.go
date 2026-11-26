package setting

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
)

type step int

const (
	stepChooseAction step = iota // Create Room vs Join Room
	stepCreateConfig             // Choose Side + Time Format
	stepJoinInput                // Enter Room Code
	stepWaiting                  // Host waiting for friend
)

// Emitted when a room is ready and both players are paired
type StartOnlineGameMsg struct {
	Room       *middleware.Room
	Player     *middleware.Player
	Assigned   middleware.Side
	TimeFormat middleware.TimeFormat
	Board      [8][8]string
}

type BackToMenuMsg struct{}

type SettingModel struct {
	width  int
	height int
	step   step

	// Shared lobby & engine hook
	lobby        *middleware.Lobby
	playerName   string
	sshPort      int
	initialBoard [8][8]string
	applyMoveFn  func([8][8]string, string) ([8][8]string, error)

	// Cursor & selections
	actionIdx   int // 0 = Create Room, 1 = Join Room
	fieldIdx    int // 0 = Side, 1 = Time Format, 2 = Start Waiting
	sideIdx     int // 0 = White, 1 = Black, 2 = Random
	timeIdx     int
	startHostFn func() error

	roomInput textinput.Model
	room      *middleware.Room
	player    *middleware.Player
	errMsg    string
}

func InitSettingModel(
	lobby *middleware.Lobby,
	playerName string,
	sshPort int,
	initialBoard [8][8]string,
	applyFn func([8][8]string, string) ([8][8]string, error),
) SettingModel {
	ti := textinput.New()
	ti.Placeholder = "e.g. 5UJ8S"
	ti.CharLimit = 5
	ti.SetWidth(16)

	return SettingModel{
		step:         stepChooseAction,
		lobby:        lobby,
		playerName:   playerName,
		sshPort:      sshPort,
		initialBoard: initialBoard,
		applyMoveFn:  applyFn,
		roomInput:    ti,
	}
}

// Add `startHostFn func() error` field to SettingModel struct
func (m SettingModel) WithStartHost(fn func() error) SettingModel {
	m.startHostFn = fn
	return m
}

func (m SettingModel) Init() tea.Cmd {
	return nil
}
