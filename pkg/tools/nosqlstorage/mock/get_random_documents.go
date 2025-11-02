package mock

import (
	"math/rand"
)

// GetRandomDocuments - Get random documents from a collection.
func (m *MockClient) GetRandomDocuments(collection string, sampleSize int) []map[string]interface{} {
	return m.doRandomRequest(collection, sampleSize, nil)
}

// GetRandomDocumentsWithFilter - Get random documents from a collection with a matcher.
func (m *MockClient) GetRandomDocumentsWithFilter(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{} {
	return m.doRandomRequest(collection, sampleSize, filters)
}

// doRandomRequest - Execute the request to the database based on all the information given in parameters.
func (m *MockClient) doRandomRequest(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{} {
	documents := m.getDataFromMockFiles(collection, nil)

	filteredDocuments := []map[string]interface{}{}
	for _, d := range documents {
		isCorrect := true
		for k, v := range filters {
			if d[k] != v {
				isCorrect = false
			}
		}
		if isCorrect {
			filteredDocuments = append(filteredDocuments, d)
		}
	}

	response := []map[string]interface{}{}
	if len(filteredDocuments) > 0 {
		for i := 1; i <= sampleSize; i++ {
			response = append(response, filteredDocuments[rand.Intn(len(filteredDocuments))])
		}
	}

	return response
}
