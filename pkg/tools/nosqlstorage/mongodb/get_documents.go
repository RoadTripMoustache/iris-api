package mongodb

import (
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetDocuments - Get all the documents from a collection.
func (m *MongoDBClient) GetDocuments(collection string, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	return m.doRequest(collection, nil, limit, offset, filters)
}

// GetDocumentsOrderBy - Get all the documents from a collection order by the `orderBy` parameter.
func (m *MongoDBClient) GetDocumentsOrderBy(collection string, orderBy bson.D, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	return m.doRequest(collection, &orderBy, limit, offset, filters)
}

// GetFirstDocument - Get the first document from a collection matching the filters.
func (m *MongoDBClient) GetFirstDocument(collection string, filters []utils.Filter) []map[string]interface{} {
	limit := 1
	offset := 0
	return m.doRequest(collection, nil, &limit, &offset, filters)
}

// doRequest - Execute the request to the database based on all the information given in parameters.
func (m *MongoDBClient) doRequest(collection string, orderBy *bson.D, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	collectionMongo := m.Client.Database(*config.GetConfigs().Database.Mongo.Name).Collection(collection)

	opts := options.Find()

	if orderBy != nil {
		opts = opts.SetSort(orderBy)
	}

	if offset != nil {
		opts = opts.SetSkip(int64(*offset))
	}

	if limit != nil {
		opts = opts.SetLimit(int64(*limit))
	}

	cur, err := collectionMongo.Find(m.Context, m.ConvertFilters(filters), opts)
	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	var mapDocs []map[string]interface{}
	for cur.Next(m.Context) {
		var document bson.M
		err = cur.Decode(&document)
		if err != nil {
			logging.Error(err, map[string]interface{}{"origin": "Decode BSON"})
		}
		mapDocs = append(mapDocs, document)
	}

	return mapDocs
}
