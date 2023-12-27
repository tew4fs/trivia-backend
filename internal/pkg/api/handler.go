package api

import (
	"encoding/json"
	"net/http"
	"tew4fs/trivia-backend/internal/pkg/action"
	"tew4fs/trivia-backend/internal/pkg/constant"
	"tew4fs/trivia-backend/internal/pkg/model"
	"tew4fs/trivia-backend/internal/pkg/response"

	"go.uber.org/zap"
)

func (a *App) HandleCurrentQuestion(w http.ResponseWriter, r *http.Request) {
	response.CreateJSONResponse(w, a.state.CurrentQuestion)
}

func (a *App) HandleNextQuestion(w http.ResponseWriter, r *http.Request) {
	a.state.NextQuestion()
	response.CreateJSONResponse(w, a.state.CurrentQuestion)
}

func (a *App) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response.CreateJSONResponse(w, model.HealthResponse{
		Status: http.StatusOK,
	})
}

func (a *App) HandleCheckAnswer(w http.ResponseWriter, r *http.Request) {
	var requestBody model.CheckAnswerRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// TODO: Add error reponse
		a.logger.Error("error checking answer", zap.Error(err))
		return
	}

	if requestBody.Question != a.state.CurrentQuestion.Question {
		a.logger.Error("trying to check question that isn't the current question")
		return
	}

	isCorrect := action.CheckAnswer(a.state.CurrentQuestion, requestBody.Answer)

	err = a.state.UpdateScore(requestBody.User, isCorrect)
	if err != nil {
		a.logger.Error("could not update score for user", zap.String(constant.UserLogKey, requestBody.User), zap.Error(err))
	}

	a.logger.Info("answer checked", zap.String("Question", requestBody.Question), zap.Bool("is correct", isCorrect))

	response.CreateJSONResponse(w, model.CheckAnswerResponse{
		User:            requestBody.User,
		Question:        requestBody.Question,
		SubmittedAnswer: requestBody.Answer,
		Correct:         isCorrect,
	})
}

func (a *App) HandleAddPlayer(w http.ResponseWriter, r *http.Request) {
	var requestBody model.AddPlayerRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		// TODO: Add error reponse
		a.logger.Error("error checking answer", zap.Error(err))
		return
	}

	err = a.state.AddPlayer(requestBody.User)
	if err != nil {
		a.logger.Error("cannot add new player", zap.String(constant.UserLogKey, requestBody.User), zap.Error(err))
		response.CreateJSONResponse(w, model.AddPlayerResponse{
			User:  requestBody.User,
			Added: false,
		})
	} else {
		response.CreateJSONResponse(w, model.AddPlayerResponse{
			User:  requestBody.User,
			Added: true,
		})
	}
}

func (a *App) HandleListScores(w http.ResponseWriter, r *http.Request) {
	response.CreateJSONResponse(w, model.CheckScoresResponse{
		Scores: a.state.PlayerScores,
	})
}
