package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	APICallsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "guide_nestor_api",
		Name:      "api_calls_counter",
		Help:      "Number of calls on the API",
	}, []string{"method", "uri", "http_code", "user"})
)
