package api

var (
	health = "/health"
)

func (a *App) setupRoutes() {
	r := a.router
	r.Use(a.LoggingMiddleware)
	r.Get(health, a.HandleHealth)
}
