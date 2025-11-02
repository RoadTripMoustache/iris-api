package images

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	imageService "github.com/RoadTripMoustache/iris_api/pkg/services/images"
)

// DeleteImage handles DELETE /v1/images/{filename}
// Removes an image file from the server storage.
func DeleteImage(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	filename := ctx.Vars["filename"]
	if filename == "" {
		return nil, errors.New(enum.BadRequest, "filename is required")
	}
	if err := imageService.DeleteImage(filename); err != nil {
		return nil, err
	}
	return utils.PrepareResponse(map[string]any{
		"deleted":  true,
		"filename": filename,
	})
}
