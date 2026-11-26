package setting

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type settingModel struct {
	pieceColor string
	timeFormat time.Time
	gameType   string
}

func (m settingModel) Init() tea.Cmd {
	return nil
}
