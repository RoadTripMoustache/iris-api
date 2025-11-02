package request

import "github.com/RoadTripMoustache/iris_api/pkg/enum"

type CreateIdeaReq struct {
	Title       string       `json:"title"`
	Tag         enum.IdeaTag `json:"tag"`
	Images      []string     `json:"images"`
	Description string       `json:"description"`
}
