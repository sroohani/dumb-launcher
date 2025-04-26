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

type userInterfaceCommand = int

const (
	uiCmdAdd userInterfaceCommand = iota
	uiCmdDelete
	uiCmdEdit
	uiCmdNone
)

type model struct {
	displayName   string
	versionString string
	width         int
	height        int
	launchables   []launchable
	cursor        string
	colorPalette  *ColorPalette
	uiCmd         userInterfaceCommand
}

func initialModel(displayName string, versionString string) model {
	return model{
		displayName:   displayName,
		versionString: versionString,
		colorPalette:  DefaultColorPalette(),
		uiCmd:         uiCmdNone,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
