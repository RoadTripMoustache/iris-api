package middlewares

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/errors"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
	"net/http"
)

func HandleResponse(data []byte, enhancedErr *errors.EnhancedError, w http.ResponseWriter, defaultHTTPCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if enhancedErr != nil {
		errArgs, err := enhancedErr.ToError()
		logging.Error(err, errArgs)
		w.WriteHeader(enhancedErr.ErrorHTTPCode)
		w.Write(enhancedErr.ToJSON())
	} else {
		w.WriteHeader(defaultHTTPCode)
		w.Write(data)
	}
}
