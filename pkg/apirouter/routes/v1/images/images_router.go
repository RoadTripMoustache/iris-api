// Package images contains all `v1/images/` API routes definitions.
package images

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/middlewares"
	imagesController "github.com/RoadTripMoustache/iris_api/pkg/controllers/images"
	"github.com/gorilla/mux"
	"net/http"
)

type ImagesRouter struct {
	MuxRouter *mux.Router
	Path      string
}

func New(router *mux.Router, basePath string) *ImagesRouter {
	return &ImagesRouter{
		MuxRouter: router,
		Path:      basePath + "/images",
	}
}

func (r *ImagesRouter) InitRoutes() {
	middlewares.AddRoute(
		r.MuxRouter,
		http.MethodPost,
		r.Path,
		http.StatusOK,
		imagesController.UploadImage,
	)
	// Get a public image by filename
	r.MuxRouter.HandleFunc(r.Path+"/{filename}", imagesController.GetImage).Methods("GET")
}
