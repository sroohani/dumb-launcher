package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type launchableVariation = int

type launchable struct {
	title     string
	variation launchableVariation
	children  []launchable
}

type model struct {
	displayName   string
	versionString string
	width         int
	height        int
	launchables   []launchable
	cursor        string
}

func initialModel(displayName string, versionString string) model {
	return model{
		displayName:   displayName,
		versionString: versionString,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
