package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	_ "github.com/charmbracelet/huh"
)

func Launch(displayName string, versionString string) {

	// contentStyle := lipgloss.NewStyle().
	// 	Width(m.width - 2).
	// 	Height(m.height - 6) // Adjust for path bar, search bar, and command bar

	// searchBarStyle := lipgloss.NewStyle().
	// 	BorderStyle(lipgloss.NormalBorder()).
	// 	BorderTop(true).
	// 	Width(m.width-2).
	// 	Padding(0, 1)

	// cmdBarStyle := lipgloss.NewStyle().
	// 	BorderStyle(lipgloss.NormalBorder()).
	// 	BorderTop(true).
	// 	Width(m.width-2).
	// 	Padding(0, 1)

	// Placeholder content

	// content := contentStyle.Render("Future list of apps and groups will appear here")

	// searchBar := searchBarStyle.Render("Search: _")

	// cmdBar := cmdBarStyle.Render("[Q] Quit  [A] Add  [D] Delete  [E] Edit")

	// Combine all sections
	p := tea.NewProgram(initialModel(displayName, versionString), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
