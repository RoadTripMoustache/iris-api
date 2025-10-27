package configs

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	configsSvc "github.com/RoadTripMoustache/iris_api/pkg/services/configs"
)

// GetConfig - Returns the images configuration used by the API.
func GetConfig(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	cfg := configsSvc.GetConfigs()
	return utils.PrepareResponse(cfg)
}
