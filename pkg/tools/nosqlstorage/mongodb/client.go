// Package mongodb contains all methods implementation to communicate with MongoDB database.
package mongodb

import (
	"context"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Client  *mongo.Client
	Context context.Context
}

// New - Generates a MongoDBClient instance from the credentials file and the project id.
func New() *MongoDBClient {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*config.GetConfigs().Database.Mongo.URI))
	if err != nil {
		logging.Fatal(err, nil)
	}

	return &MongoDBClient{
		Client:  client,
		Context: ctx,
	}
}
