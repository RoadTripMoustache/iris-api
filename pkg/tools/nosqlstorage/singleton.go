// Package nosqlstorage contains all the code to communicate with NoSQL databases.
package nosqlstorage

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/mock"
	mongodb "github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage/mongodb"
)

var noSQLStorageInstance NoSQLClient = nil

// GetInstance - Get the instance of the datasource client.
func GetInstance() NoSQLClient {
	if noSQLStorageInstance == nil {
		if config.GetConfigs().Database.Mock.Enabled {
			noSQLStorageInstance = mock.New()
		} else if config.GetConfigs().Database.Mongo.URI != nil {
			noSQLStorageInstance = mongodb.New()
		}
	}
	return noSQLStorageInstance
}
