package response

import "github.com/RoadTripMoustache/iris_api/pkg/enum"

type ListIdeaResponse struct {
	ID           string       `json:"id"`
	CreatedAt    interface{}  `json:"created_at"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	Tag          enum.IdeaTag `json:"tag"`
	CreatorID    string       `json:"creator_id"`
	VotesCount   int          `json:"votes_count"`
	IsOpen       bool         `json:"is_open"`
	UserHasVoted bool         `json:"user_has_voted,omitempty"`
}
