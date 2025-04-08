package main

import (
	"log"
	"log/slog"
	"rps/api"
	"rps/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Panic("error loading the config")
	}
	slog.Info("loaded config", "config", cfg)

	api.InitRoutes()
}
