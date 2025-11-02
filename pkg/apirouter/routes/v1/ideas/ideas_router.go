// Package ideas contains all `v1/ideas/` API routes definitions.
package ideas

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	ctrl "github.com/RoadTripMoustache/iris_api/pkg/controllers/ideas"
	"github.com/gorilla/mux"
	"net/http"
)

// IdeasRouter handles routing for ideas-related endpoints.
type IdeasRouter struct {
	MuxRouter *mux.Router
	Path      string
}

// New creates a new IdeasRouter instance.
func New(router *mux.Router, parentPath string) *IdeasRouter {
	return &IdeasRouter{
		MuxRouter: router,
		Path:      fmt.Sprintf("%s/ideas", parentPath),
	}
}

// InitRoutes initializes the routes for the ideas API.
func (p *IdeasRouter) InitRoutes() {
	// List ideas
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodGet,
		p.Path,
		http.StatusOK,
		ctrl.ListIdeas,
	)
	// Get one idea (full details)
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodGet,
		p.Path+"/{id}",
		http.StatusOK,
		ctrl.GetIdea,
	)
	// Create idea
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPost,
		p.Path,
		http.StatusCreated,
		ctrl.CreateIdea,
	)
	// Vote
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPost,
		p.Path+"/{id}/vote",
		http.StatusOK,
		ctrl.VoteIdea,
	)
	// Unvote
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodDelete,
		p.Path+"/{id}/vote",
		http.StatusOK,
		ctrl.UnvoteIdea,
	)
	// Add comment
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPost,
		p.Path+"/{id}/comments",
		http.StatusCreated,
		ctrl.AddComment,
	)
	// Edit comment
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPatch,
		p.Path+"/{id}/comments/{commentId}",
		http.StatusOK,
		ctrl.EditComment,
	)
	// Delete comment
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodDelete,
		p.Path+"/{id}/comments/{commentId}",
		http.StatusOK,
		ctrl.DeleteComment,
	)
	// Open/Close idea (admin)
	middlewares.AddRoute(
		p.MuxRouter,
		http.MethodPatch,
		p.Path+"/{id}/open",
		http.StatusOK,
		ctrl.AdminSetIdeaOpen,
	)
}
