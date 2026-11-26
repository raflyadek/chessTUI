package setting

import (
	"fmt"
	"net"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/raflyadek/chessTUI/internal/middleware"
	"github.com/raflyadek/chessTUI/internal/terminal/style"
)

func (m SettingModel) View() string {
	var body string
	switch m.step {
	case stepChooseAction:
		opts := []string{"Create Room", "Join Room"}
		var lines []string
		lines = append(lines, style.SectionTitleStyle.Render("PLAY WITH FRIENDS (ONLINE)"))
		for i, opt := range opts {
			prefix := "  "
			st := style.NormalStyle
			if i == m.actionIdx {
				prefix = "> "
				st = style.SelectedStyle
			}
			lines = append(lines, st.Render(prefix+opt))
		}
		lines = append(lines, "", style.SecondaryStyle.Render("j/k: Move  Enter: Select  Esc: Back"))
		body = lipgloss.JoinVertical(lipgloss.Left, lines...)

	case stepCreateConfig:
		sides := []string{"White", "Black", "Random"}
		tf := middleware.TimePresets[m.timeIdx].Label
		rows := []string{
			style.SectionTitleStyle.Render("ROOM SETTINGS"),
			m.renderRow(0, fmt.Sprintf("Side        : < %s >", sides[m.sideIdx])),
			m.renderRow(1, fmt.Sprintf("Time Format : < %s >", tf)),
			"",
			m.renderRow(2, "[ Create & Wait for Friend ]"),
			"",
			style.SecondaryStyle.Render("h/l: Change  j/k: Next  Enter: Confirm  Esc: Back"),
		}
		body = lipgloss.JoinVertical(lipgloss.Left, rows...)

	case stepJoinInput:
		lines := []string{
			style.SectionTitleStyle.Render("JOIN ROOM"),
			style.SecondaryStyle.Render("Enter the 5-character room code from your friend:"),
			style.InputStyle.Render(m.roomInput.View()),
		}
		if m.errMsg != "" {
			lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("#E05A47")).Render(m.errMsg))
		}
		lines = append(lines, "", style.SecondaryStyle.Render("Enter: Join  Esc: Back"))
		body = lipgloss.JoinVertical(lipgloss.Left, lines...)

	case stepWaiting:
		lanIP := localIPv4()
		sshCmd := fmt.Sprintf("ssh %s -p %d", lanIP, m.sshPort)
		lines := []string{
			style.SectionTitleStyle.Render("WAITING FOR FRIEND..."),
			// fmt.Sprintf("Room Code : %s", style.HeaderStyle.Render(m.room.ID)),
			fmt.Sprintf("Your Side : %s", strings.ToUpper(string(m.player.Side))),
			fmt.Sprintf("Format    : %s", m.room.TimeFormat.Label),
			"",
			style.SecondaryStyle.Render("Ask your friend on the same network to run:"),
			style.HeaderStyle.Render(sshCmd),
			style.SecondaryStyle.Render("Then choose Play With Friends -> Join Room and enter the code."),
			"",
			style.SecondaryStyle.Render("Esc: Cancel Room"),
		}
		body = lipgloss.JoinVertical(lipgloss.Left, lines...)
	}

	card := style.PanelStyle.Padding(1, 3).Render(body)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, card)
}

func (m SettingModel) renderRow(idx int, label string) string {
	if m.fieldIdx == idx {
		return style.SelectedStyle.Render("> " + label)
	}
	return style.NormalStyle.Render("  " + label)
}

func waitForPlayerEvent(ch <-chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return <-ch
	}
}

func localIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "<your-ip>"
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return "<your-ip>"
}
