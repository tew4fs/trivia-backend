package api

import (
	"fmt"
	"net/http"
	"tew4fs/trivia-backend/internal/pkg/constant"

	"go.uber.org/zap"
)

func (a *App) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info(fmt.Sprintf("endpoint hit: %s", r.URL),
			zap.String(constant.MethodLogKey, r.Method),
			zap.String(constant.UrlLogKey, r.URL.String()),
			zap.String(constant.ProtocolLogKey, r.Proto),
		)
		next.ServeHTTP(w, r)
	})
}
