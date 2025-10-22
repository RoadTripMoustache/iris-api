package api

import (
	"github.com/RoadTripMoustache/grostube/pkg/apirouter/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateLinks(t *testing.T) {
	sizeList := 12
	pageNumber := 2
	firstPageNumber := 1
	lastPageNumber := 4
	pageSize := 3

	testCases := []struct {
		caseDesc         string
		baseURL          string
		p                utils.Pagination
		sizeList         *int
		otherQueryParams string
		expectedResult   Links
	}{
		{
			caseDesc:         "Full empty",
			baseURL:          "toto.com",
			p:                utils.Pagination{},
			sizeList:         nil,
			otherQueryParams: "",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=0",
				Previous: "toto.com?page[number]=1&page[size]=0",
				Self:     "toto.com?page[number]=1&page[size]=0",
				Next:     "toto.com?page[number]=1&page[size]=0",
				Last:     "toto.com?page[number]=1&page[size]=0",
				HasMore:  false,
			},
		},
		{
			caseDesc:         "Full empty with query param",
			baseURL:          "toto.com",
			p:                utils.Pagination{},
			sizeList:         nil,
			otherQueryParams: "queryParam",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=0",
				Previous: "toto.com?page[number]=1&page[size]=0",
				Self:     "toto.com?page[number]=1&page[size]=0",
				Next:     "toto.com?page[number]=1&page[size]=0",
				Last:     "toto.com?page[number]=1&page[size]=0",
				HasMore:  false,
			},
		},
		{
			caseDesc: "Pagination filled with query param",
			baseURL:  "toto.com",
			p: utils.Pagination{
				PageNumber: &pageNumber,
				PageSize:   &pageSize,
			},
			sizeList:         &sizeList,
			otherQueryParams: "queryParam",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=3",
				Previous: "toto.com?page[number]=1&page[size]=3",
				Self:     "toto.com?page[number]=2&page[size]=3",
				Next:     "toto.com?page[number]=3&page[size]=3",
				Last:     "toto.com?page[number]=4&page[size]=3",
				HasMore:  true,
			},
		},
		{
			caseDesc: "Pagination filled with query param",
			baseURL:  "toto.com",
			p: utils.Pagination{
				PageNumber: &pageNumber,
				PageSize:   &pageSize,
			},
			sizeList:         &sizeList,
			otherQueryParams: "queryParam",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=3",
				Previous: "toto.com?page[number]=1&page[size]=3",
				Self:     "toto.com?page[number]=2&page[size]=3",
				Next:     "toto.com?page[number]=3&page[size]=3",
				Last:     "toto.com?page[number]=4&page[size]=3",
				HasMore:  true,
			},
		},
		{
			caseDesc: "First page",
			baseURL:  "toto.com",
			p: utils.Pagination{
				PageNumber: &firstPageNumber,
				PageSize:   &pageSize,
			},
			sizeList:         &sizeList,
			otherQueryParams: "queryParam",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=3",
				Previous: "toto.com?page[number]=1&page[size]=3",
				Self:     "toto.com?page[number]=1&page[size]=3",
				Next:     "toto.com?page[number]=2&page[size]=3",
				Last:     "toto.com?page[number]=4&page[size]=3",
				HasMore:  true,
			},
		},
		{
			caseDesc: "Last page",
			baseURL:  "toto.com",
			p: utils.Pagination{
				PageNumber: &lastPageNumber,
				PageSize:   &pageSize,
			},
			sizeList:         &sizeList,
			otherQueryParams: "queryParam",
			expectedResult: Links{
				First:    "toto.com?page[number]=1&page[size]=3",
				Previous: "toto.com?page[number]=3&page[size]=3",
				Self:     "toto.com?page[number]=4&page[size]=3",
				Next:     "toto.com?page[number]=4&page[size]=3",
				Last:     "toto.com?page[number]=4&page[size]=3",
				HasMore:  false,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.caseDesc, func(t *testing.T) {
			ctx := utils.Context{
				Pagination: testCase.p,
				RequestURI: testCase.baseURL,
			}
			result := GenerateLinks(ctx, testCase.sizeList)

			assert.NotNil(t, result)
			assert.Equal(t, result, testCase.expectedResult)
		})
	}
}
