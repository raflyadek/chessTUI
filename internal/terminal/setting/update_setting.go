package setting

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
)

func (m SettingModel) Update(msg tea.Msg) (SettingModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case middleware.MatchStartedMsg:
		room, ok := m.lobby.GetRoom(msg.RoomID)
		if !ok {
			room = m.room
		}
		return m, func() tea.Msg {
			return StartOnlineGameMsg{
				Room:       room,
				Player:     m.player,
				Assigned:   msg.Assigned,
				TimeFormat: msg.TimeFormat,
				Board:      msg.Board,
			}
		}

	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc":
			if m.step == stepWaiting && m.room != nil {
				m.lobby.RemoveRoom(m.room.ID)
				m.room = nil
			}
			if m.step == stepChooseAction {
				return m, func() tea.Msg { return BackToMenuMsg{} }
			}
			m.step = stepChooseAction
			m.errMsg = ""
			return m, nil
		}

		switch m.step {
		case stepChooseAction:
			return m.updateChooseAction(msg.String())
		case stepCreateConfig:
			return m.updateCreateConfig(msg.String())
		case stepJoinInput:
			return m.updateJoinInput(msg)
		}
	}

	return m, nil
}

func (m SettingModel) updateChooseAction(key string) (SettingModel, tea.Cmd) {
	switch key {
	case "j", "down":
		m.actionIdx = (m.actionIdx + 1) % len(m.opts)
	case "k", "up":
		m.actionIdx = (m.actionIdx - 1 + len(m.opts)) % len(m.opts)
	case "enter":
		m.errMsg = ""
		if m.actionIdx == 0 {
			m.step = stepCreateConfig
			m.waiting = waitingLocal
			m.fieldIdx = 0
		} else if m.actionIdx == 1 {
			//TODO:create new step for online??
			m.step = stepCreateConfig
			m.waiting = waitingOnline
			m.fieldIdx = 0
			m.isOnline = true
		} else if m.actionIdx == 2 {
			m.step = stepJoinInput
			m.roomInput.Reset()
			m.roomInput.Focus()
			return m, textinput.Blink
		}
	}
	return m, nil
}

func (m SettingModel) updateCreateConfig(key string) (SettingModel, tea.Cmd) {
	sides := []middleware.Side{middleware.SideWhite, middleware.SideBlack, middleware.SideRandom}

	switch key {
	case "j", "down":
		m.fieldIdx = (m.fieldIdx + 1) % 3
	case "k", "up":
		m.fieldIdx = (m.fieldIdx + 2) % 3
	case "h", "left":
		if m.fieldIdx == 0 {
			m.sideIdx = (m.sideIdx + len(sides) - 1) % len(sides)
		} else if m.fieldIdx == 1 {
			m.timeIdx = (m.timeIdx + len(middleware.TimePresets) - 1) % len(middleware.TimePresets)
		}
	case "l", "right":
		if m.fieldIdx == 0 {
			m.sideIdx = (m.sideIdx + 1) % len(sides)
		} else if m.fieldIdx == 1 {
			m.timeIdx = (m.timeIdx + 1) % len(middleware.TimePresets)
		}
	case "enter":
		if m.fieldIdx < 3 {
			m.fieldIdx++
			return m, nil
		}

		//start ssh server local on demand
		if m.startHostFn != nil {
			if err := m.startHostFn(); err != nil {
				m.errMsg = fmt.Sprintf("Failed to start SSH host: %v", err)
				return m, nil
			}
		}

		room, player := m.lobby.CreateRoom(
			m.playerName,
			sides[m.sideIdx],
			middleware.TimePresets[m.timeIdx],
			m.initialBoard,
			m.applyMoveFn,
		)
		m.room = room
		m.player = player
		m.step = stepWaiting
		return m, waitForPlayerEvent(player.Events)
	}
	return m, nil
}

func (m SettingModel) updateJoinInput(msg tea.KeyPressMsg) (SettingModel, tea.Cmd) {
	if msg.String() == "enter" {
		code := strings.ToUpper(strings.TrimSpace(m.roomInput.Value()))
		if len(code) != 5 {
			m.errMsg = "Enter a 5-character room code"
			return m, nil
		}
		room, player, err := m.lobby.JoinRoom(code, m.playerName)
		if err != nil {
			m.errMsg = err.Error()
			return m, nil
		}
		m.room = room
		m.player = player
		// Guest immediately receives MatchStartedMsg from JoinRoom
		return m, waitForPlayerEvent(player.Events)
	}

	var cmd tea.Cmd
	m.roomInput, cmd = m.roomInput.Update(msg)
	return m, cmd
}
