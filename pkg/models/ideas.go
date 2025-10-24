package models

import (
	"time"
)

type IdeaTag string

const (
	IdeaTagBug         IdeaTag = "bug"
	IdeaTagEnhancement IdeaTag = "enhancement"
)

// Comment represents a comment on an idea with an internal id used by the API
type Comment struct {
	ID        string    `json:"id" bson:"id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Message   string    `json:"message" bson:"message"`
	Images    []string  `json:"images" bson:"images"`
}

// Idea represents an idea. It has a Mongo internal _id and an application-level id
// used for all operations (get, vote, manage comments).
type Idea struct {
	ID           string    `json:"id" bson:"id"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	Title        string    `json:"title" bson:"title"`
	Description  string    `json:"description" bson:"description"`
	Tag          IdeaTag   `json:"tag" bson:"tag"`
	CreatorID    string    `json:"creator_id" bson:"creator_id"`
	VotesCount   int       `json:"votes_count" bson:"votes_count"`
	Voters       []string  `json:"voters" bson:"voters"`
	Comments     []Comment `json:"comments" bson:"comments"`
	Images       []string  `json:"images" bson:"images"`
	IsOpen       bool      `json:"is_open" bson:"is_open"`
	UserHasVoted bool      `json:"user_has_voted,omitempty" bson:"-"`
}

const (
	MaxImagesPerEntity = 5
	MaxImageSizeBytes  = 2 * 1024 * 1024 // 2MB
)
