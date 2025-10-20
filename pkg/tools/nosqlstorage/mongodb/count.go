package mongodb

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/utils"
)

// Count - Returns the number of elements in the collection given in parameter.
func (m *MongoDBClient) Count(collection string, filters []utils.Filter) *int {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)

	int64Count, err := collectionMongo.CountDocuments(m.Context, m.ConvertFilters(filters))
	if err != nil {
		logging.Error(err, nil)
		return nil
	}
	intCount := int(int64Count)
	return &intCount
}
