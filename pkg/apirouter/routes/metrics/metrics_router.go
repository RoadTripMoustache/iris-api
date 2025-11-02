// Package metrics contains all `metrics/` API routes definitions.
package metrics

import (
	"github.com/RoadTripMoustache/iris_api/pkg/monitoring/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type APIMetricsRouter struct {
	MuxRouter *mux.Router
	Path      string
}

// New - Generate a new APIMetricsRouter
func New(router *mux.Router) *APIMetricsRouter {
	return &APIMetricsRouter{
		MuxRouter: router,
		Path:      "/metrics",
	}
}

// InitRoutes - Initialize all the routes for the "metrics" path.
func (a *APIMetricsRouter) InitRoutes() {
	prometheus.MustRegister(metrics.APICallsDurationHistogram)

	a.MuxRouter.Handle(a.Path, promhttp.Handler()).Methods(http.MethodGet)
}
