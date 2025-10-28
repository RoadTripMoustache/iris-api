package dbmodels

import (
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"time"
)

const (
	IdeaCollectionName string = "ideas"
	IdeaIDColumn       string = "id"
)

// Idea represents an idea. It has a Mongo internal _id and an application-level id
// used for all operations (get, vote, manage comments).
type Idea struct {
	ID           string       `json:"id" bson:"id"`
	CreatedAt    time.Time    `json:"created_at" bson:"created_at"`
	Title        string       `json:"title" bson:"title"`
	Description  string       `json:"description" bson:"description"`
	Tag          enum.IdeaTag `json:"tag" bson:"tag"`
	CreatorID    string       `json:"creator_id" bson:"creator_id"`
	VotesCount   int          `json:"votes_count" bson:"votes_count"`
	Voters       []string     `json:"voters" bson:"voters"`
	Comments     []Comment    `json:"comments" bson:"comments"`
	Images       []string     `json:"images" bson:"images"`
	IsOpen       bool         `json:"is_open" bson:"is_open"`
	UserHasVoted bool         `json:"user_has_voted,omitempty" bson:"-"`
}
