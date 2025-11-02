// Package apirouter contains all routes definitions and middlewares.
package apirouter

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/routes/metrics"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type MetricsRouter struct {
	MuxRouter *mux.Router
}

// NewMetricsRouter - Generate a new MetricsRouter
func NewMetricsRouter() *MetricsRouter {
	return &MetricsRouter{
		MuxRouter: mux.NewRouter(),
	}
}

// Serve - Init all the subrouters, middlewares and start the server.
func (a *MetricsRouter) Serve() {
	// Init routes
	metrics.New(a.MuxRouter).InitRoutes()

	// Start the HTTP Server
	logging.Info("Prometheus exposure started", nil)
	http.ListenAndServe(config.GetConfigs().Server.MetricsExpose, cors.AllowAll().Handler(a.MuxRouter))
}
