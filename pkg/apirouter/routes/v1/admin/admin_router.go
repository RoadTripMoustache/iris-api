// Package admin contains all `v1/admin/` API routes definitions.
package admin

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	ctrl "github.com/RoadTripMoustache/iris_api/pkg/controllers/admin"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	adminsvc "github.com/RoadTripMoustache/iris_api/pkg/services/admin"
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
	// Get the admin list
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodGet,
		p.Path,
		http.StatusOK,
		func(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
			adm, e := adminsvc.GetAdmin(ctx, ctx.UserEmail)
			if e != nil || adm == nil {
				return nil, errors.New(enum.AuthUnauthorized, nil)
			}

			return ctrl.GetAdmins(ctx)
		},
	)
	// Add an user as an admin
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPost,
		p.Path,
		http.StatusOK,
		func(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
			adm, e := adminsvc.GetAdmin(ctx, ctx.UserEmail)
			if e != nil || adm == nil {
				return nil, errors.New(enum.AuthUnauthorized, nil)
			}

			return ctrl.AddAdmin(ctx)
		},
	)
	// Remove an user as an admin
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodDelete,
		p.Path,
		http.StatusOK,
		func(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
			adm, e := adminsvc.GetAdmin(ctx, ctx.UserEmail)
			if e != nil || adm == nil {
				return nil, errors.New(enum.AuthUnauthorized, nil)
			}

			return ctrl.RemoveAdmin(ctx)
		},
	)
	// Check if current user is admin
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodGet,
		p.Path+"/me",
		http.StatusOK,
		ctrl.IsCurrentUserAdmin,
	)
}
