package mongodb

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

// Delete - Delete an element in the selected collection
func (m *MongoDBClient) Delete(collection string, idValue interface{}, idParamLabel string) error {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)
	_, err := collectionMongo.DeleteOne(m.Context, bson.M{idParamLabel: idValue})
	return err
}
