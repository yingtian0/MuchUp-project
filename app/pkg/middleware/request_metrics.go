package middleware

import (
	"net/http"
	"time"

	"MuchUp/app/pkg/logger"

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
			rec := &statusRecorder{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next.ServeHTTP(rec, r)

			elapsed := time.Since(start)
			path := r.URL.Path

			appLogger.Infof(
				"[request] method=%s path=%s status=%d latency_ms=%d",
				r.Method,
				path,
				rec.status,
				elapsed.Milliseconds(),
			)
		})
	}
}
