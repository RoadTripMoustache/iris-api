// Package ideas contains all the methods of the IdeaService
package ideas

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"strings"
	"time"

	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/dbmodels"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	appErrors "github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	nosql "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collection = "ideas"

func CreateIdea(ctx apiUtils.Context, title string, description string, tag enum.IdeaTag, creatorID string, images []string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	if title == "" || (tag != enum.IdeaTagBug && tag != enum.IdeaTagEnhancement) {
		return nil, appErrors.New(enum.BadRequest, "invalid title or tag")
	}
	appConfig := config.GetConfigs()
	if len(images) > appConfig.Images.MaxImagesPerIdea {
		return nil, appErrors.New(enum.TooManyImages, "too many images")
	}
	idea := &dbmodels.Idea{
		ID:          primitive.NewObjectID().Hex(),
		CreatedAt:   time.Now().UTC(),
		Title:       title,
		Description: description,
		Tag:         tag,
		CreatorID:   creatorID,
		VotesCount:  0,
		Voters:      []string{},
		Comments:    []dbmodels.Comment{},
		Images:      images,
		IsOpen:      true,
	}
	if err := nosql.GetInstance().Add(collection, idea); err != nil {
		logging.Error(err, map[string]interface{}{"service": "ideas", "method": "CreateIdea"})
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return idea, nil
}

func ListIdeas(ctx apiUtils.Context) ([]dbmodels.Idea, *appErrors.EnhancedError) {
	docs := nosql.GetInstance().GetDocumentsOrderBy(collection, bson.D{{Key: "created_at", Value: -1}}, ctx.Pagination.PageSize, ctx.Pagination.GetOffset(), []nosqlUtils.Filter{})
	ideas := make([]dbmodels.Idea, 0, len(docs))
	for _, d := range docs {
		b, _ := bson.Marshal(d)
		var idea dbmodels.Idea
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

// GetIdea returns a single idea by id with user_has_voted computed
func GetIdea(ctx apiUtils.Context, ideaID string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
	_ = bson.Unmarshal(b, &idea)
	if idea.Voters != nil && ctx.UserID != "" {
		for _, v := range idea.Voters {
			if v == ctx.UserID {
				idea.UserHasVoted = true
				break
			}
		}
	}
	return &idea, nil
}

func Vote(ctx apiUtils.Context, ideaID string, userID string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
	_ = bson.Unmarshal(b, &idea)
	for _, v := range idea.Voters {
		if v == userID {
			return &idea, nil
		}
	}
	idea.Voters = append(idea.Voters, userID)
	idea.VotesCount = len(idea.Voters)
	if err := nosql.GetInstance().Update(collection, idea.ID, "id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

func Unvote(ctx apiUtils.Context, ideaID string, userID string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
	_ = bson.Unmarshal(b, &idea)
	newVoters := make([]string, 0, len(idea.Voters))
	for _, v := range idea.Voters {
		if v != userID {
			newVoters = append(newVoters, v)
		}
	}
	idea.Voters = newVoters
	idea.VotesCount = len(idea.Voters)
	if err := nosql.GetInstance().Update(collection, idea.ID, "id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

func AddComment(ctx apiUtils.Context, ideaID string, userID string, message string, images []string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	if strings.TrimSpace(message) == "" {
		return nil, appErrors.New(enum.BadRequest, "message required")
	}
	appConfig := config.GetConfigs()
	if len(images) > appConfig.Images.MaxImagesPerComment {
		return nil, appErrors.New(enum.TooManyImages, "too many images")
	}
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
	_ = bson.Unmarshal(b, &idea)
	comment := dbmodels.Comment{
		ID:        primitive.NewObjectID().Hex(),
		CreatedAt: time.Now().UTC(),
		UserID:    userID,
		Message:   message,
		Images:    images,
	}
	idea.Comments = append(idea.Comments, comment)
	if err := nosql.GetInstance().Update(collection, idea.ID, "id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// SetIdeaOpen updates the open state of an idea
func SetIdeaOpen(ctx apiUtils.Context, ideaID string, isOpen bool) (*dbmodels.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
	_ = bson.Unmarshal(b, &idea)
	idea.IsOpen = isOpen
	if err := nosql.GetInstance().Update(collection, idea.ID, "_id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// EditComment updates a user's own comment on an idea
func EditComment(ctx apiUtils.Context, ideaID, commentID, userID, message string, images []string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	appConfig := config.GetConfigs()
	if strings.TrimSpace(message) == "" {
		return nil, appErrors.New(enum.BadRequest, "message required")
	}
	if len(images) > appConfig.Images.MaxImagesPerComment {
		return nil, appErrors.New(enum.TooManyImages, "too many images")
	}
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
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
	if err := nosql.GetInstance().Update(collection, idea.ID, "id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}

// DeleteComment removes a user's own comment from an idea
func DeleteComment(ctx apiUtils.Context, ideaID, commentID, userID string) (*dbmodels.Idea, *appErrors.EnhancedError) {
	doc := nosql.GetInstance().GetFirstDocument(collection, []nosqlUtils.Filter{{Param: "id", Value: ideaID, Operator: "eq"}})
	if len(doc) == 0 {
		return nil, appErrors.New(enum.ResourceNotFound, "idea")
	}
	b, _ := bson.Marshal(doc[0])
	var idea dbmodels.Idea
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
	if err := nosql.GetInstance().Update(collection, idea.ID, "id", idea); err != nil {
		return nil, appErrors.New(enum.InternalServerError, err)
	}
	return &idea, nil
}
