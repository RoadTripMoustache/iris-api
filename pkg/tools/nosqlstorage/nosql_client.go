package nosqlstorage

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type NoSQLClient interface {
	// GetDocuments - Get all the documents from a collection.
	GetDocuments(collection string, limit *int, offset *int, filters []utils.Filter) []map[string]interface{}

	// GetFirstDocument - Get the first document from a collection matching the filters.
	GetFirstDocument(collection string, filters []utils.Filter) []map[string]interface{}

	// GetDocumentsOrderBy - Get all the documents from a collection order by the `orderBy` parameter.
	GetDocumentsOrderBy(collection string, orderBy bson.D, limit *int, offset *int, filters []utils.Filter) []map[string]interface{}

	// GetRandomDocuments - Get random documents from a collection.
	GetRandomDocuments(collection string, sampleSize int) []map[string]interface{}

	// GetRandomDocumentsWithFilter - Get random documents from a collection with a matcher.
	GetRandomDocumentsWithFilter(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{}

	// Count - Returns the number of elements in the collection given in parameter.
	Count(collection string, filters []utils.Filter) *int

	// Add - Add data in the selected collection
	Add(collection string, data interface{}) error

	// Delete - Delete an element in the selected collection
	Delete(collection string, idValue interface{}, idParamLabel string) error

	// Update - Update data in the selected collection
	Update(collection string, idValue interface{}, idLabel string, data interface{}) error
}
