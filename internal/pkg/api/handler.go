package api

import (
	"net/http"
	"tew4fs/golang-api-skeleton/internal/pkg/model"
	"tew4fs/golang-api-skeleton/internal/pkg/response"
)

func (a *App) HandleHealth(w http.ResponseWriter, r *http.Request) {
	response.CreateJSONResponse(w, model.HealthResponse{
		Status: http.StatusOK,
	})
}
