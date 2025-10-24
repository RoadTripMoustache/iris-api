package ideas

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models/api/request"
	"github.com/RoadTripMoustache/iris_api/pkg/services/ideas"
	"strings"
)

func CreateIdea(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := request.CreateIdeaReq{}
	if err := utils.BodyFormatter(ctx.Body, &payload); err != nil {
		return nil, errors.New(enum.BadRequest, err)
	}
	if strings.TrimSpace(payload.Description) == "" {
		return nil, errors.New(enum.BadRequest, "description required")
	}
	if e := validateImages(payload.Images); e != nil {
		return nil, e
	}
	idea, e := ideas.CreateIdea(ctx, payload.Title, payload.Description, payload.Tag, ctx.UserID, payload.Images)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}
