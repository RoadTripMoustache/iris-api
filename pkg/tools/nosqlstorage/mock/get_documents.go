package mock

import (
	"fmt"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

// GetDocuments - Get all the documents from a collection.
func (m *MockClient) GetDocuments(collection string, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	documents := m.getDataFromMockFiles(collection, filters)

	if offset != nil {
		if *offset > len(documents) {
			return make([]map[string]interface{}, 0)
		}
		documents = documents[*offset:]
	}

	if limit != nil && len(documents) > *limit {
		documents = documents[:*limit]
	}

	return documents
}

// GetDocumentsOrderBy - Get all the documents from a collection order by the `orderBy` parameter.
func (m *MockClient) GetDocumentsOrderBy(collection string, orderBy bson.D, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	documents := m.getDataFromMockFiles(collection, filters)

	sort.Slice(documents, func(i, j int) bool {
		for _, e := range orderBy {
			if fmt.Sprintf("%v", documents[i][e.Key]) != fmt.Sprintf("%v", documents[j][e.Key]) {
				if e.Value == 1 {
					return fmt.Sprintf("%v", documents[i][e.Key]) < fmt.Sprintf("%v", documents[j][e.Key])
				} else {
					return fmt.Sprintf("%v", documents[i][e.Key]) > fmt.Sprintf("%v", documents[j][e.Key])
				}
			}
		}
		return true
	})

	if offset != nil {
		if *offset > len(documents) {
			return make([]map[string]interface{}, 0)
		}
		documents = documents[*offset:]
	}

	if limit != nil && len(documents) > *limit {
		documents = documents[:*limit]
	}
	return documents
}

// GetFirstDocument - Get the first document from a collection matching the filters.
func (m *MockClient) GetFirstDocument(collection string, filters []utils.Filter) []map[string]interface{} {
	documents := m.getDataFromMockFiles(collection, filters)
	if len(documents) == 0 {
		return nil
	}
	return []map[string]interface{}{documents[0]}
}
