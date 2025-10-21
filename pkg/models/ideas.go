package models

import (
	"time"
)

type IdeaTag string

const (
	IdeaTagBug         IdeaTag = "bug"
	IdeaTagEnhancement IdeaTag = "enhancement"
)

type Comment struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Message   string    `json:"message" bson:"message"`
	Images    []string  `json:"images" bson:"images"`
}

type Idea struct {
	ID           string    `json:"id" bson:"_id,omitempty"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	Title        string    `json:"title" bson:"title"`
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
