package api

var (
	health          = "/health"
	currentQuestion = "/question/current"
	nextQuestion    = "/question/next"
	checkAnswer     = "/question/answer"
	addPlayer       = "/players/add"
	listScores      = "/players/scores"
)

func (a *App) setupRoutes() {
	r := a.router
	r.Use(a.LoggingMiddleware)
	r.Get(health, a.HandleHealth)
	r.Get(currentQuestion, a.HandleCurrentQuestion)
	r.Get(nextQuestion, a.HandleNextQuestion)
	r.Post(checkAnswer, a.HandleCheckAnswer)
	r.Post(addPlayer, a.HandleAddPlayer)
	r.Get(listScores, a.HandleListScores)
}
