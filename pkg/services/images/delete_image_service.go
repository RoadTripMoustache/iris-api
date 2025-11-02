package images

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	appErrors "github.com/RoadTripMoustache/iris_api/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
)

var filenameSafeRegexpSvc = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)

// DeleteImage removes an image file from the server storage by its filename.
func DeleteImage(filename string) *appErrors.EnhancedError {
	if filename == "" || !filenameSafeRegexpSvc.MatchString(filename) {
		return appErrors.New(enum.BadRequest, "invalid filename")
	}

	cfg := config.GetConfigs().Server
	fullPath := filepath.Join(cfg.ImagesDir, filename)

	if _, err := os.Stat(fullPath); err != nil {
		if os.IsNotExist(err) {
			return appErrors.New(enum.ResourceNotFound, "image")
		}
		return appErrors.New(enum.InternalServerError, err)
	}

	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return appErrors.New(enum.ResourceNotFound, "image")
		}
		return appErrors.New(enum.InternalServerError, err)
	}
	return nil
}
