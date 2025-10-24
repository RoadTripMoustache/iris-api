package ideas

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/services/ideas"
)

func GetIdea(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	ideaID := ctx.Vars["id"]
	idea, e := ideas.GetIdea(ctx, ideaID)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}
