package mock

import (
	utils2 "github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/utils"
)

// Count - Returns the number of elements in the collection given in parameter.
func (m *MockClient) Count(collection string, filters []utils2.Filter) *int {
	i := len(m.getDataFromMockFiles(collection, filters))
	return &i
}
