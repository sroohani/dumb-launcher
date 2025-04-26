package tui

import (
	"fmt"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	title := fmt.Sprintf("%v %v ", m.displayName, m.versionString)
	titleBarWidth := utf8.RuneCountInString(title)

	titleBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(m.colorPalette.TitleBarBg)).
		Foreground(lipgloss.Color(m.colorPalette.TitleBarFg)).
		Width(titleBarWidth).
		Padding(0, 0, 0, 0)
	titleBar := titleBarStyle.Render(title)

	searchBarWidth := m.width - titleBarWidth
	searchBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(m.colorPalette.SearchBarBg)).
		Foreground(lipgloss.Color(m.colorPalette.SearchBarFg)).
		Width(searchBarWidth).
		Padding(0, 0, 0, 2)
	searchBar := searchBarStyle.Render("[Start typing to search...]")

	contentStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(m.colorPalette.ContentBg)).
		Foreground(lipgloss.Color(m.colorPalette.ContentFg)).
		Width(m.width).
		Height(m.height - 3)
	content := contentStyle.Render("Future list of apps and groups will appear here")

	cmdBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(m.colorPalette.CmdBarBg)).
		Foreground(lipgloss.Color(m.colorPalette.CmdBarFg)).
		Width(m.width - 1)

	underlineStyle := lipgloss.NewStyle().
		Underline(false).
		Background(lipgloss.Color(m.colorPalette.CmdBarFg)).
		Foreground(lipgloss.Color(m.colorPalette.CmdBarBg))
	cmdRanges := []struct {
		cmd string
		rng lipgloss.Range
	}{
		{
			cmd: "Add ",
			rng: lipgloss.NewRange(0, 1, underlineStyle),
		},
		{
			cmd: "Delete ",
			rng: lipgloss.NewRange(0, 1, underlineStyle),
		},

		{
			cmd: "Edit ",
			rng: lipgloss.NewRange(0, 1, underlineStyle),
		},
		{
			cmd: "Quit",
			rng: lipgloss.NewRange(0, 1, underlineStyle),
		},
	}
	cmdBar := "Ctrl + ["
	for _, cr := range cmdRanges {
		cmdBar += lipgloss.StyleRanges(cr.cmd, cr.rng)
	}
	cmdBar += "]"
	cmdBarStyle.Render(cmdBar)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, titleBar, searchBar),
		content,
		cmdBar,
	)
}
