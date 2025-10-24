package ideas

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models/api/response"
	"github.com/RoadTripMoustache/iris_api/pkg/services/ideas"
	nosql "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
)

func ListIdeas(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	list, e := ideas.ListIdeas(ctx)
	if e != nil {
		return nil, e
	}
	out := make([]response.ListIdeaResponse, 0, len(list))
	for _, it := range list {
		out = append(out, response.ListIdeaResponse{
			ID:           it.ID,
			CreatedAt:    it.CreatedAt,
			Title:        it.Title,
			Description:  it.Description,
			Tag:          it.Tag,
			CreatorID:    it.CreatorID,
			VotesCount:   it.VotesCount,
			IsOpen:       it.IsOpen,
			UserHasVoted: it.UserHasVoted,
		})
	}
	// total number of items in the resource (across all pages)
	countPtr := nosql.GetInstance().Count("ideas", []nosqlUtils.Filter{})
	total := 0
	if countPtr != nil {
		total = *countPtr
	}
	return utils.PrepareListResponse(ctx, out, total)
}
