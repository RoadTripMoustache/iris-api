package mongodb

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

// Update - Update data in the selected collection
func (m *MongoDBClient) Update(collection string, idValue interface{}, idLabel string, data interface{}) error {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)

	singleResult := collectionMongo.FindOneAndReplace(m.Context, bson.D{{idLabel, idValue}}, data)
	return singleResult.Err()
}
