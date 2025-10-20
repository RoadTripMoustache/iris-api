package middlewares

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/errors"
	"github.com/gorilla/mux"
	"net/http"
)

func AddRoute(
	r *mux.Router,
	method string,
	path string,
	defaultHTTPCode int,
	f func(utils.Context) ([]byte, *errors.EnhancedError),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {

		ctx := GenerateContext(req, method, path)
		data, err := f(ctx)
		HandleResponse(data, err, w, defaultHTTPCode)
	}).Methods(method)
}

func AddSSERoute(
	r *mux.Router,
	method string,
	path string,
	f func(utils.Context, http.ResponseWriter, *http.ResponseController, <-chan struct{}),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		_, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Le serveur ne supporte pas le SSE", http.StatusInternalServerError)
			return
		}

		// Set http headers required for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create a channel for client disconnection
		clientGone := req.Context().Done()

		// Create response controller
		rc := http.NewResponseController(w)

		ctx := GenerateContext(req, method, path)
		f(ctx, w, rc, clientGone)
	}).Methods(method)
}
