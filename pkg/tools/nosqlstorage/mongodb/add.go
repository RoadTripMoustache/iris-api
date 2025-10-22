package mongodb

import "github.com/RoadTripMoustache/iris_api/pkg/config"

// Add - Add data in the selected collection
func (m *MongoDBClient) Add(collection string, data interface{}) error {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)

	_, err := collectionMongo.InsertOne(m.Context, data)

	return err
}
