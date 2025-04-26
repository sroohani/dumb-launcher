package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	_ "github.com/charmbracelet/huh"
)

func Launch(displayName string, versionString string) {
	p := tea.NewProgram(initialModel(displayName, versionString), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
