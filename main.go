package main

import (
	"github.com/sroohani/dumb-launcher/cfg"
	"github.com/sroohani/dumb-launcher/tui"
)

func main() {
	tui.Launch(displayName, versionString)
	var cfgParams cfg.Cfg
	cfg.Read(&cfgParams)
}
