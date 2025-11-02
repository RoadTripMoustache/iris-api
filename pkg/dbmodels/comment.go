// Package dbmodels contains all the DB models
package dbmodels

import (
	"time"
)

// Comment represents a comment on an idea with an internal id used by the API
type Comment struct {
	ID        string    `json:"id" bson:"id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Message   string    `json:"message" bson:"message"`
	Images    []string  `json:"images" bson:"images"`
}
