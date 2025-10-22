package ideas

import (
	"path/filepath"
	"strings"

	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/validators"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models"
	"github.com/RoadTripMoustache/iris_api/pkg/services/ideas"
	nosql "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
)

func validateImages(urls []string) *errors.EnhancedError {
	if len(urls) > models.MaxImagesPerEntity {
		return errors.New(enum.BadRequest, "too many images")
	}
	for _, u := range urls {
		ext := strings.ToLower(filepath.Ext(u))
		if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
			return errors.New(enum.BadRequest, "invalid image type; only .png or .jpg allowed")
		}
	}
	return nil
}

type createIdeaReq struct {
	Title       string         `json:"title"`
	Tag         models.IdeaTag `json:"tag"`
	Images      []string       `json:"images"`
	Description string         `json:"description"`
}

type addCommentReq struct {
	Message string   `json:"message"`
	Images  []string `json:"images"`
}

type setOpenReq struct {
	IsOpen bool `json:"is_open"`
}

func EditComment(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := addCommentReq{}
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
	commentID := ctx.Vars["commentId"]
	idea, e := ideas.EditComment(ctx, ideaID, commentID, ctx.UserID, payload.Message, payload.Images)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}

func DeleteComment(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	ideaID := ctx.Vars["id"]
	commentID := ctx.Vars["commentId"]
	idea, e := ideas.DeleteComment(ctx, ideaID, commentID, ctx.UserID)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}

func CreateIdea(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := createIdeaReq{}
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

func ListIdeas(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	list, e := ideas.ListIdeas(ctx)
	if e != nil {
		return nil, e
	}
	// total number of items in the resource (across all pages)
	countPtr := nosql.GetInstance().Count("ideas", []nosqlUtils.Filter{})
	total := 0
	if countPtr != nil {
		total = *countPtr
	}
	return utils.PrepareListResponse(ctx, list, total)
}

func VoteIdea(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	ideaID := ctx.Vars["id"]
	idea, e := ideas.Vote(ctx, ideaID, ctx.UserID)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}

func UnvoteIdea(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	ideaID := ctx.Vars["id"]
	idea, e := ideas.Unvote(ctx, ideaID, ctx.UserID)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(idea)
}

func AddComment(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := addCommentReq{}
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

// AdminSetIdeaOpen allows an admin to open/close an idea
func AdminSetIdeaOpen(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	if err := validators.IsAdmin(ctx); err != nil {
		return nil, err
	}
	payload := setOpenReq{}
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
