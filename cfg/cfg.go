package cfg

import (
	"log"

	"github.com/spf13/viper"
)

type Cfg struct {
	DbUrl string
}

func Read(cfg *Cfg) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg.DbUrl = viper.Get("DUMB_LAUNCHER_DB_URL").(string)
}
