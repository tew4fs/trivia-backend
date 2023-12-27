package api

import (
	"fmt"
	"net/http"
	"tew4fs/trivia-backend/internal/pkg/config"
	"tew4fs/trivia-backend/internal/pkg/state"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type APIService interface {
	Start()
}

type App struct {
	cfg    config.AppConfig
	logger *zap.Logger
	server http.Server
	router *chi.Mux
	state  *state.State
}

func NewApp(cfg config.AppConfig, logger *zap.Logger) *App {
	s := &App{
		cfg:    cfg,
		logger: logger,
		router: chi.NewRouter(),
		state:  state.NewState(),
	}

	s.server = http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.AppHost, cfg.Port),
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.setupRoutes()

	return s
}

func (a *App) Start() {
	a.logger.Info(fmt.Sprintf("Starting server at address %s", a.server.Addr))
	a.logger.Fatal("API shutting down", zap.Error(a.server.ListenAndServe()))
}
