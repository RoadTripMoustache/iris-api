// Package v1 contains all `v1/` API routes definitions.
package v1

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/routes/v1/ideas"
	"github.com/gorilla/mux"
)

type APIV1Router struct {
	MuxRouter *mux.Router
	Path      string
}

// New - Generate a new APIV1Router
func New(router *mux.Router) *APIV1Router {
	return &APIV1Router{
		MuxRouter: router,
		Path:      "/v1",
	}
}

// InitRoutes - Initialize all the routes for the "v1" path.
func (a *APIV1Router) InitRoutes() {
	ideas.New(a.MuxRouter, a.Path).InitRoutes()
}
