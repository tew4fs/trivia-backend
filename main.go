package main

import (
	"tew4fs/golang-api-skeleton/internal/pkg/api"
	"tew4fs/golang-api-skeleton/internal/pkg/config"
	"tew4fs/golang-api-skeleton/internal/pkg/log"
)

func main() {
	cfg := config.LoadConfigs()
	logger := log.GetLogger(cfg)

	app := api.NewApp(cfg, logger)

	app.Start()

}
