package ideas

import (
	"strings"
	"time"

	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	appErrors "github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	nosql "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collection = "ideas"

func CreateIdea(ctx apiUtils.Context, title string, description string, tag models.IdeaTag, creatorID string, images []string) (*models.Idea, *appErrors.EnhancedError) {
	if title == "" || (tag != models.IdeaTagBug && tag != models.IdeaTagEnhancement) {
		return nil, appErrors.New(enum.BadRequest, "invalid title or tag")
	}
	if len(images) > models.MaxImagesPerEntity {
		return nil, appErrors.New(enum.BadRequest, "too many images")
	}
	idea := &models.Idea{
		CreatedAt:   time.Now().UTC(),
		Title:       title,
		Description: description,
		Tag:         tag,
		CreatorID:   creatorID,
		VotesCount:  0,
		Voters:      []string{},
		Comments:    []models.Comment{},
		Images:      images,
		IsOpen:      true,
	}
	if err := nosql.GetInstance().Add(collection, idea); err != nil {
		logging.Error(err, map[string]interface{}{"service": "ideas", "method": "CreateIdea"})
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return idea, nil
}

func ListIdeas(ctx apiUtils.Context) ([]models.Idea, *appErrors.EnhancedError) {
	docs := nosql.GetInstance().GetDocumentsOrderBy(collection, bson.D{{Key: "created_at", Value: -1}}, ctx.Pagination.PageSize, ctx.Pagination.GetOffset(), []nosqlUtils.Filter{})
	ideas := make([]models.Idea, 0, len(docs))
	for _, d := range docs {
		b, _ := bson.Marshal(d)
		var idea models.Idea
		_ = bson.Unmarshal(b, &idea)
		if idea.Voters != nil && ctx.UserID != "" {
			for _, v := range idea.Voters {
				if v == ctx.UserID {
					idea.UserHasVoted = true
					break
				}
			}
		}
		ideas = append(ideas, idea)
	}
	return ideas, nil
}

func Vote(ctx apiUtils.Context, ideaID string, userID string) (*models.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	for _, v := range idea.Voters {
		if v == userID {
			return &idea, nil
		}
	}
	idea.Voters = append(idea.Voters, userID)
	idea.VotesCount = len(idea.Voters)
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"voters": idea.Voters, "votes_count": idea.VotesCount}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

func Unvote(ctx apiUtils.Context, ideaID string, userID string) (*models.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	newVoters := make([]string, 0, len(idea.Voters))
	for _, v := range idea.Voters {
		if v != userID {
			newVoters = append(newVoters, v)
		}
	}
	idea.Voters = newVoters
	idea.VotesCount = len(idea.Voters)
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"voters": idea.Voters, "votes_count": idea.VotesCount}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

func AddComment(ctx apiUtils.Context, ideaID string, userID string, message string, images []string) (*models.Idea, *appErrors.EnhancedError) {
	if strings.TrimSpace(message) == "" {
		return nil, appErrors.New(enum.BadRequest, "message required")
	}
	if len(images) > models.MaxImagesPerEntity {
		return nil, appErrors.New(enum.BadRequest, "too many images")
	}
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	comment := models.Comment{
		ID:        primitive.NewObjectID().Hex(),
		CreatedAt: time.Now().UTC(),
		UserID:    userID,
		Message:   message,
		Images:    images,
	}
	idea.Comments = append(idea.Comments, comment)
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"comments": idea.Comments}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// SetIdeaOpen updates the open state of an idea
func SetIdeaOpen(ctx apiUtils.Context, ideaID string, isOpen bool) (*models.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	idea.IsOpen = isOpen
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"is_open": idea.IsOpen}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// EditComment updates a user's own comment on an idea
func EditComment(ctx apiUtils.Context, ideaID, commentID, userID, message string, images []string) (*models.Idea, *appErrors.EnhancedError) {
	if strings.TrimSpace(message) == "" {
		return nil, appErrors.New(enum.BadRequest, "message required")
	}
	if len(images) > models.MaxImagesPerEntity {
		return nil, appErrors.New(enum.BadRequest, "too many images")
	}
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	found := false
	for i, c := range idea.Comments {
		if c.ID == commentID {
			if c.UserID != userID {
				return nil, appErrors.New(enum.AuthUnauthorized, "cannot modify this comment")
			}
			idea.Comments[i].Message = message
			idea.Comments[i].Images = images
			found = true
			break
		}
	}
	if !found {
		return nil, appErrors.New(enum.ResourceNotFound, "comment")
	}
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"comments": idea.Comments}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// DeleteComment removes a user's own comment from an idea
func DeleteComment(ctx apiUtils.Context, ideaID, commentID, userID string) (*models.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "_id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea models.Idea
	_ = bson.Unmarshal(b, &idea)
	idx := -1
	for i, c := range idea.Comments {
		if c.ID == commentID {
			if c.UserID != userID {
				return nil, appErrors.New(enum.AuthUnauthorized, "cannot delete this comment")
			}
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, appErrors.New(enum.ResourceNotFound, "comment")
	}
	idea.Comments = append(idea.Comments[:idx], idea.Comments[idx+1:]...)
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", map[string]interface{}{"comments": idea.Comments}); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}
