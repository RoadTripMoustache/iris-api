package api

import (
	"fmt"
	"github.com/RoadTripMoustache/grostube/pkg/apirouter/utils"
	"math"
	"strings"
)

type Links struct {
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Self     string `json:"self,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
	HasMore  bool   `json:"has_more"`
}

func GenerateLinks(ctx utils.Context, sizeList *int) Links {
	baseURL := strings.Split(ctx.RequestURI, "?")[0]
	otherQueryParams := ""
	for k, v := range ctx.QueryParams {
		if k != "page[number]" && k != "page[size]" {
			otherQueryParams = fmt.Sprintf("%s&%s=%s", otherQueryParams, k, strings.Join(v, ","))
		}
	}

	if sizeList == nil {
		defaultSizeList := 0
		sizeList = &defaultSizeList
	}

	if ctx.Pagination.PageSize == nil {
		ctx.Pagination.PageSize = sizeList
	}
	if ctx.Pagination.PageNumber == nil {
		firstPage := 1
		ctx.Pagination.PageNumber = &firstPage
	}

	maxPage := 1
	if ctx.Pagination.PageSize != nil && sizeList != nil && *sizeList > 0 {
		maxPage = int(math.Ceil(float64(*sizeList) / float64(*ctx.Pagination.PageSize)))
	}

	links := Links{}

	links.First = fmt.Sprintf("%s?page[number]=1&page[size]=%d%s", baseURL, *ctx.Pagination.PageSize, otherQueryParams)

	previousPage := *ctx.Pagination.PageNumber - 1
	if sizeList != nil && previousPage > *sizeList {
		previousPage = *sizeList
	}
	if previousPage < 1 {
		previousPage = 1
	}
	links.Previous = fmt.Sprintf("%s?page[number]=%d&page[size]=%d%s", baseURL, previousPage, *ctx.Pagination.PageSize, otherQueryParams)

	links.Self = fmt.Sprintf("%s?page[number]=%d&page[size]=%d%s", baseURL, *ctx.Pagination.PageNumber, *ctx.Pagination.PageSize, otherQueryParams)

	nextPage := *ctx.Pagination.PageNumber + 1
	if nextPage > maxPage {
		nextPage = maxPage
	}
	links.Next = fmt.Sprintf("%s?page[number]=%d&page[size]=%d%s", baseURL, nextPage, *ctx.Pagination.PageSize, otherQueryParams)
	links.Last = fmt.Sprintf("%s?page[number]=%d&page[size]=%d%s", baseURL, maxPage, *ctx.Pagination.PageSize, otherQueryParams)

	links.HasMore = *ctx.Pagination.PageNumber < maxPage

	return links
}
