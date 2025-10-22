package utils

import (
	"encoding/json"
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models/api"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
)

func PrepareResponse(a any) ([]byte, *errors.EnhancedError) {
	responseData, err := json.Marshal(a)
	if err != nil {
		logging.Error(err, nil)
		return nil, errors.New(enum.BadRequest, nil)
	}
	return responseData, nil
}

// PrepareListResponse wraps list data into the standard api.Response with pagination links.
// Data contains the payload currently returned by the endpoint.
// sizeList should be the total number of items in the resource (across all pages).
func PrepareListResponse(ctx apiUtils.Context, data any, sizeList int) ([]byte, *errors.EnhancedError) {
	resp := api.Response{
		Links: api.GenerateLinks(ctx, &sizeList),
		Data:  data,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		logging.Error(err, nil)
		return nil, errors.New(enum.BadRequest, nil)
	}
	return b, nil
}
