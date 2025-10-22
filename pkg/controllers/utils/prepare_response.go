package utils

import (
	"encoding/json"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
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
