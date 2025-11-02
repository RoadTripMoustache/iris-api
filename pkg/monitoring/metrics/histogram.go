package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	APICallsDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "iris_api",
			Name:      "api_requests_duration_seconds",
			Help:      "Duration of HTTP requests in seconds.",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		},
		[]string{"method", "uri", "http_code"},
	)
)
