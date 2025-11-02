package mongodb

import (
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func (m *MongoDBClient) ConvertFilters(filters []utils.Filter) interface{} {
	// TODO : Faire une enumération des opérateurs mis en place pour les filtres
	mongoFilters := bson.D{}
	for _, filter := range filters {
		if filter.Operator == "eq" || filter.Operator == "==" {
			mongoFilters = append(mongoFilters, bson.E{Key: filter.Param, Value: filter.Value})
		} else if filter.Operator == "in" {
			mongoFilters = append(mongoFilters, bson.E{Key: filter.Param, Value: bson.D{{"$in", filter.Value}}})
		} else if filter.Operator == "array-contains" {
			mongoFilters = append(mongoFilters, bson.E{Key: filter.Param, Value: bson.D{{"$all", bson.A{filter.Value}}}})
		} else if filter.Operator == "or" {
			return bson.M{"$or": filter.Value.([]bson.M)}
		} else if filter.Operator == "date-gt" {
			mongoFilters = append(mongoFilters, bson.E{Key: filter.Param, Value: bson.M{"$gt": filter.Value.(time.Time)}})
		}
	}

	return mongoFilters
}
