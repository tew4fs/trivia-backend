package api

import (
	"fmt"
	"net/http"
	"tew4fs/golang-api-skeleton/internal/pkg/config"
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
}

func NewApp(cfg config.AppConfig, logger *zap.Logger) *App {
	s := &App{
		cfg:    cfg,
		logger: logger,
		router: chi.NewRouter(),
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
	a.logger.Info(fmt.Sprintf("Starting server on port %d", a.cfg.Port))
	a.logger.Fatal("API shutting down", zap.Error(a.server.ListenAndServe()))
}
