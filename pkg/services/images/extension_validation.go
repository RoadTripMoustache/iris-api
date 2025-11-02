package images

import (
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"path/filepath"
	"strings"
)

func ExtensionValidation(url string) *errors.EnhancedError {
	appConfig := config.GetConfigs()
	ext := strings.ToLower(filepath.Ext(url))
	isExtensionAllowed := false
	for _, extension := range appConfig.Images.AcceptedExtensions {
		if ext == fmt.Sprintf(".%s", extension) {
			isExtensionAllowed = true
		}
	}
	if !isExtensionAllowed {
		return errors.New(enum.ImageExtensionNotAllowed, "Image extension not allowed")
	}
	return nil
}
