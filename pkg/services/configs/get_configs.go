package configs

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	resp "github.com/RoadTripMoustache/iris_api/pkg/models/api/response"
)

// GetConfigs builds the configs to share to the front.
func GetConfigs() resp.Configurations {
	cfg := config.GetConfigs().Images
	return resp.Configurations{
		MaxImagesPerIdea:    cfg.MaxImagesPerIdea,
		MaxImagesPerComment: cfg.MaxImagesPerComment,
		MaxSize:             cfg.MaxSize,
		AcceptedExtensions:  cfg.AcceptedExtensions,
	}
}
