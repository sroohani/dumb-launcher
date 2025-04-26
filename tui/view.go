package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	pathBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#00ff00")).
		Foreground(lipgloss.Color("#000000")).
		Width(m.width-2).
		Padding(0, 1)
	pathBar := pathBarStyle.Render("Path: / (Home)")

	searchBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#0000ff")).
		Foreground(lipgloss.Color("#ffffff")).
		Width(m.width-2).
		Padding(0, 1)
	searchBar := searchBarStyle.Render("Search: _")

	contentStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#000000")).
		Foreground(lipgloss.Color("#ffffff")).
		Width(m.width - 2).
		Height(m.height - 3)
	content := contentStyle.Render("Future list of apps and groups will appear here")

	cmdBarStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#ffffff")).
		Foreground(lipgloss.Color("#000000")).
		Width(m.width-2).
		Padding(0, 1)
	cmdBar := cmdBarStyle.Render("[Q] Quit  [A] Add  [D] Delete  [E] Edit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		pathBar,
		searchBar,
		content,
		cmdBar,
	)
}
