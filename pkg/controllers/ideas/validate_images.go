package ideas

import (
	"github.com/RoadTripMoustache/iris_api/pkg/constantes"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"path/filepath"
	"strings"
)

func validateImages(urls []string) *errors.EnhancedError {
	if len(urls) > constantes.MaxImagesPerEntity {
		return errors.New(enum.TooManyImages, "too many images")
	}
	for _, u := range urls {
		ext := strings.ToLower(filepath.Ext(u))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			return errors.New(enum.BadRequest, "invalid image type; only .png or .jpg allowed")
		}
	}
	return nil
}
