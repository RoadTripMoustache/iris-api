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

func AddComment(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := request.AddCommentReq{}
	if err := utils.BodyFormatter(ctx.Body, &payload); err != nil {
		return nil, errors.New(enum.BadRequest, err)
	}
	if strings.TrimSpace(payload.Message) == "" {
		return nil, errors.New(enum.BadRequest, "message required")
	}
	if e := validateImages(payload.Images); e != nil {
		return nil, e
	}
	ideaID := ctx.Vars["id"]
	idea, e := ideas.AddComment(ctx, ideaID, ctx.UserID, payload.Message, payload.Images)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}
