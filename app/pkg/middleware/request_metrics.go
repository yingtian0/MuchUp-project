package middleware

import (
	"MuchUp/app/pkg/logger"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func RequestMetrics(appLogger logger.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			next.ServeHTTP(w, r)
			elapsed := time.Since(start)
			rec := &statusRecorder{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			path := r.Response.Request.URL.Path

			appLogger.Info(
				"[request] method=%s path=%s status=%d latency_ms=%d",
				r.Method,
				path,
				rec.status,
				elapsed.Milliseconds(),
			)
		})

	}
}
