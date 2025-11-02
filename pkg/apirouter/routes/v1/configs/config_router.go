// Package configs contains all `v1/configs/` API routes definitions.
package configs

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	configsController "github.com/RoadTripMoustache/iris_api/pkg/controllers/configs"
	"github.com/gorilla/mux"
	"net/http"
)

type ConfigRouter struct {
	MuxRouter *mux.Router
	Path      string
}

func New(router *mux.Router, basePath string) *ConfigRouter {
	return &ConfigRouter{
		MuxRouter: router,
		Path:      basePath + "/configs",
	}
}

func (r *ConfigRouter) InitRoutes() {
	// Get configuration
	middlewares.AddRoute(
		r.MuxRouter,
		http.MethodGet,
		r.Path,
		http.StatusOK,
		configsController.GetConfig,
	)
}
