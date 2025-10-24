package ideas

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/validators"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models/api/request"
	"github.com/RoadTripMoustache/iris_api/pkg/services/ideas"
)

// AdminSetIdeaOpen allows an admin to open/close an idea
func AdminSetIdeaOpen(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	if err := validators.IsAdmin(ctx); err != nil {
		return nil, err
	}
	payload := request.SetOpenReq{}
	if err := utils.BodyFormatter(ctx.Body, &payload); err != nil {
		return nil, errors.New(enum.BadRequest, err)
	}
	ideaID := ctx.Vars["id"]
	idea, e := ideas.SetIdeaOpen(ctx, ideaID, payload.IsOpen)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}
