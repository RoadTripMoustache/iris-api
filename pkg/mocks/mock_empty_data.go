package mocks

import (
	utils2 "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
)

var (
	EmptyEnhancedError *errors.EnhancedError = nil
	EmptyInt           *int
	EmptyString        *string
	EmptyContext                      = utils2.Context{}
	EmptyFilters       []utils.Filter = nil
)
