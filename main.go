package main

import (
	"github.com/sroohani/dumb-launcher/cfg"
	"github.com/sroohani/dumb-launcher/db"
	"github.com/sroohani/dumb-launcher/tui"
)

func main() {
	var cfgParams cfg.Cfg
	cfg.Read(&cfgParams)
	db.Load(cfgParams.DbUrl)

	tui.Launch(displayName, versionString)
}
