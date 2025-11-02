// Package admin contains all `v1/admin/` API routes definitions.
package admin

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	ctrl "github.com/RoadTripMoustache/iris_api/pkg/controllers/admin"
	"github.com/gorilla/mux"
	"net/http"
)

// AdminRouter handles routing for admin-related endpoints.
type AdminRouter struct {
	MuxRouter *mux.Router
	Path      string
}

// New creates a new AdminRouter instance.
func New(router *mux.Router, parentPath string) *AdminRouter {
	return &AdminRouter{
		MuxRouter: router,
		Path:      fmt.Sprintf("%s/admin", parentPath),
	}
}

// InitRoutes initializes the routes for the admin API.
func (p *AdminRouter) InitRoutes() {
	// Check if current user is admin
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodGet,
		p.Path+"/me",
		http.StatusOK,
		ctrl.IsCurrentUserAdmin,
	)
}
