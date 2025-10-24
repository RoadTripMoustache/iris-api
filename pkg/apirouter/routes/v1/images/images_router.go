// Package images contains all `v1/images/` API routes definitions.
package images

import (
	imagesController "github.com/RoadTripMoustache/iris_api/pkg/controllers/images"
	"github.com/gorilla/mux"
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
	r.MuxRouter.HandleFunc(r.Path, imagesController.UploadImage).Methods("POST")
}
