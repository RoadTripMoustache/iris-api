// Package apirouter contains all routes definitions and middlewares.
package apirouter

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/routes/v1"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type APIRouter struct {
	MuxRouter *mux.Router
}

// New - Generate a new APIRouter
func New() *APIRouter {
	return &APIRouter{
		MuxRouter: mux.NewRouter(),
	}
}

// Serve - Init all the subrouters, middlewares and start the server.
func (a *APIRouter) Serve() {
	// Init routes
	v1.New(a.MuxRouter).InitRoutes()

	// Declare all the middlewares
	// --- After - Must be in reverse order !
	a.MuxRouter.Use(middlewares.AfterMetricsMiddleware)
	a.MuxRouter.Use(middlewares.AfterLoggingMiddleware)
	// --- Before
	a.MuxRouter.Use(middlewares.BeforeLoggingMiddleware)
	a.MuxRouter.Use(middlewares.AuthenticationMiddleware)
	a.MuxRouter.Use(middlewares.PaginationMiddleware)
	a.MuxRouter.Use(middlewares.VersionFilterMiddleware)

	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"Authorization", "*"},
		Debug:          false, // TODO MAx
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowOriginVaryRequestFunc: func(r *http.Request, origin string) (bool, []string) {
			originsAllowed := config.GetConfigs().Server.OriginsAllowed
			isAllowed := false
			for _, originAllowed := range originsAllowed {
				if originAllowed == "*" {
					isAllowed = true
				}
				if !isAllowed && r.Header.Get("Origin") == originAllowed {
					isAllowed = true
				}
			}

			return isAllowed, []string{"*"}
		},
	}).Handler(a.MuxRouter)

	// Start the HTTP Server
	err := http.ListenAndServe(config.GetConfigs().Server.Expose, handler)
	if err != nil {
		logging.Error(err, nil)
	}
	logging.Info("API started", nil)
}
