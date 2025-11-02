package middlewares

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"net/http"
)

// BeforeLoggingMiddleware - Log every request to know every calls going in.
func BeforeLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.Info("Request received", map[string]interface{}{
			"uri":    r.RequestURI,
			"method": r.Method,
		})

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// AfterLoggingMiddleware - Log every request response to know every calls going in.
func AfterLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, _ := w.(http.Flusher)
		rec := utils.StatusRecorder{ResponseWriter: w, Status: 200, Flusher: f}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(&rec, r)

		logging.Info("Request response", map[string]interface{}{
			"uri":         r.RequestURI,
			"method":      r.Method,
			"status_code": rec.Status,
		})

	})
}
