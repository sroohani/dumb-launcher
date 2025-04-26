package tui

type ColorPalette struct {
	TitleBarBg  string
	TitleBarFg  string
	PathBarBg   string
	PathBarFg   string
	SearchBarBg string
	SearchBarFg string
	ContentBg   string
	ContentFg   string
	CmdBarBg    string
	CmdBarFg    string
}

// Gruv-Box color palette
// https://www.color-hex.com/color-palette/1026676
func DefaultColorPalette() *ColorPalette {
	return &ColorPalette{
		TitleBarBg:  "#cc241d",
		TitleBarFg:  "#3c3836",
		PathBarBg:   "#cc241d",
		PathBarFg:   "#3c3836",
		SearchBarBg: "#458588",
		SearchBarFg: "#d79921",
		ContentBg:   "#8ec07c",
		ContentFg:   "#3c3836",
		CmdBarBg:    "#d79921",
		CmdBarFg:    "#cc241d",
	}
}
