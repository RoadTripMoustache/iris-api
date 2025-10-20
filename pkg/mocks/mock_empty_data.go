package mocks

import (
	utils2 "github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/errors"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/utils"
)

var (
	EmptyEnhancedError *errors.EnhancedError = nil
	EmptyInt           *int
	EmptyString        *string
	EmptyContext                      = utils2.Context{}
	EmptyFilters       []utils.Filter = nil
)
