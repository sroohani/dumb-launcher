package tui

import (
	"fmt"
)

func (m model) View() string {
	s := fmt.Sprintf("Welcome to %v %v!\n", m.displayName, m.versionString)

	return s
}
