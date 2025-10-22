package mongodb

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRandomDocuments - Get random documents from a collection.
func (m *MongoDBClient) GetRandomDocuments(collection string, sampleSize int) []map[string]interface{} {
	return m.doRandomRequest(collection, sampleSize, nil)
}

// GetRandomDocumentsWithFilter - Get random documents from a collection with a matcher.
func (m *MongoDBClient) GetRandomDocumentsWithFilter(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{} {
	return m.doRandomRequest(collection, sampleSize, filters)
}

// doRandomRequest - Execute the request to the database based on all the information given in parameters.
func (m *MongoDBClient) doRandomRequest(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{} {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)

	pipeline := []map[string]interface{}{}
	if filters != nil {
		pipeline = append(pipeline, map[string]interface{}{
			"$match": filters,
		})
	}

	pipeline = append(pipeline, map[string]interface{}{
		"$sample": map[string]interface{}{
			"size": sampleSize,
		},
	})

	cur, err := collectionMongo.Aggregate(m.Context, pipeline)
	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	var mapDocs []map[string]interface{}
	for cur.Next(m.Context) {
		var document bson.M
		err = cur.Decode(&document)
		if err != nil {
			logging.Error(err, map[string]interface{}{"origin": "Decode BSON - doRandomRequest"})
		}
		mapDocs = append(mapDocs, document)
	}

	return mapDocs
}
