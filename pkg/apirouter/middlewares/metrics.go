package middlewares

import (
	auth2 "firebase.google.com/go/auth"
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/monitoring/metrics"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/auth"
	"net/http"
	"strings"
	"time"
)

// AfterMetricsMiddleware - Does stats on every request response to know every calls going in.
func AfterMetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, _ := w.(http.Flusher)
		rec := utils.StatusRecorder{ResponseWriter: w, Status: 200, Flusher: f}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(&rec, r)
		startTime := time.Now()

		authorizationHeader := r.Header.Get(enum.Authorization)
		if len(authorizationHeader) > 8 {
			authorizationHeader = authorizationHeader[7:]
		}
		userID := "not_logged_user"
		var token *auth2.Token = nil
		if !strings.HasPrefix(r.URL.Path, "/v1/stripe/notify") {
			token = auth.GetInstance().VerifyIDToken(authorizationHeader)
		}
		if token != nil {
			userID = token.Subject
		}

		uri := strings.Split(r.RequestURI, "?")[0]

		metrics.APICallsCounter.WithLabelValues(
			r.Method,
			uri,
			fmt.Sprintf("%d", rec.Status),
			userID,
		).Inc()

		duration := time.Since(startTime).Seconds()
		metrics.APICallsDurationHistogram.WithLabelValues(
			r.Method,
			uri,
			fmt.Sprintf("%d", rec.Status),
		).Observe(duration)
	})
}
