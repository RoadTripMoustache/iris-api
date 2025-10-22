// Package mock - Contains all the nosqlstorage methods to mock a database.
package mock

import (
	"encoding/json"
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	"github.com/RoadTripMoustache/iris_api/pkg/dbmodels/admin"
	nosqlUtils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/exp/slices"
	"os"
	"time"
)

// getDataFromMockFiles - Get the data from mock files stores in the project
// Params
// - collection : string - name of the collection which we want to get the data. Will be the name of the file to open.
// - filter : []nosqlUtils.Filter - filters to apply on the data to get.
// Returns
// - []map[string]interface{} - The filtered list of objects.
func (m *MockClient) getDataFromMockFiles(collection string, filters []nosqlUtils.Filter) []map[string]interface{} {
	filePath := *config.GetConfigs().Database.Mock.DataFolderPath + "/" + collection + ".json"

	if _, err := os.Stat(filePath); err == nil {
		// Open the json file
		f, err := os.Open(filePath)
		if err != nil {
			utils.ProcessError(err)
		}
		// Prepare to close the file at the end
		defer f.Close()

		var list = make([]map[string]interface{}, 0)

		decoder := json.NewDecoder(f)
		err = decoder.Decode(&list)
		if err != nil {
			utils.ProcessError(err)
		}

		// Insert & replace cached data
		idLabel := getIDLabel(collection)
		for _, cachedItem := range m.Cache[collection] {
			isReplaced := false
			for i, element := range list {
				if element[idLabel] == cachedItem[idLabel] {
					list[i] = cachedItem
					isReplaced = true
				}
			}
			if !isReplaced {
				list = append(list, cachedItem)
			}
		}

		// Apply filters on the retrieved data
		for _, filter := range filters {
			var filteredList = make([]map[string]interface{}, 0)
			for _, element := range list {
				if (filter.Operator == "eq" || filter.Operator == "==") && fmt.Sprintf("%v", element[filter.Param]) == fmt.Sprintf("%v", filter.Value) {
					filteredList = append(filteredList, element)
				} else if filter.Operator == "in" && slices.Contains(filter.Value.([]string), fmt.Sprintf("%v", element[filter.Param])) {
					filteredList = append(filteredList, element)
				} else if filter.Operator == "array-contains" && element[filter.Param] != nil && slices.Contains(element[filter.Param].([]interface{}), filter.Value) {
					filteredList = append(filteredList, element)
				} else if filter.Operator == "or" {
					isAdded := false
					for _, conditions := range filter.Value.([]bson.M) {
						isConditionsRespected := true
						for k, v := range conditions {
							if element[k] != v {
								isConditionsRespected = false
							}
						}
						if !isAdded && isConditionsRespected {
							filteredList = append(filteredList, element)
							isAdded = true
						}
					}
				} else if filter.Operator == "date-gt" {
					lastUpdate, _ := time.Parse(time.RFC3339, element[filter.Param].(string))
					lastUpdate = time.Date(lastUpdate.Year(), lastUpdate.Month(), lastUpdate.Day(), 0, 0, 0, 0, time.UTC)
					if lastUpdate.After(filter.Value.(time.Time)) {
						filteredList = append(filteredList, element)
					}
				}
			}
			list = filteredList
		}
		return list
	}

	return nil
}

// getIDLabel - Get the id label of a collection.
func getIDLabel(collection string) string {
	switch collection {
	case admin.AdminCollectionName:
		return admin.AdminUserIdLabel
	}
	return "uuid"
}
