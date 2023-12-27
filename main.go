package main

import (
	"tew4fs/trivia-backend/internal/pkg/api"
	"tew4fs/trivia-backend/internal/pkg/config"
	"tew4fs/trivia-backend/internal/pkg/log"
)

func main() {
	config.LoadConfigs()
	logger := log.GetLogger()

	app := api.NewApp(*config.Config, logger)

	app.Start()

}
