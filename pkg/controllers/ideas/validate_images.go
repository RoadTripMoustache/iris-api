package ideas

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/services/images"
)

func validateImages(urls []string) *errors.EnhancedError {
	appConfig := config.GetConfigs()
	if len(urls) > appConfig.Images.MaxImagesPerIdea {
		return errors.New(enum.TooManyImages, "too many images")
	}
	for _, u := range urls {
		err := images.ExtensionValidation(u)
		if err != nil {
			return err
		}
	}
	return nil
}
