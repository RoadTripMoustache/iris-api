package mock

import (
	"encoding/json"
	"fmt"
)

// Update - Update data in the selected collection
func (m *MockClient) Update(collection string, idValue interface{}, idLabel string, data interface{}) error {
	if m.Cache[collection] == nil {
		m.Cache[collection] = make([]map[string]interface{}, 0)
	}

	jsonBytes, _ := json.Marshal(data)
	var resultMap map[string]interface{}
	json.Unmarshal(jsonBytes, &resultMap)

	newCollectionList := make([]map[string]interface{}, 0)

	notInDB := true
	println(fmt.Sprintf("====== UPDATE : %s =======", idValue))
	println(fmt.Sprintf("data : %v", data))
	for _, collectionItem := range m.Cache[collection] {
		if fmt.Sprintf("%v", collectionItem[idLabel]) == idValue {
			println(fmt.Sprintf("REPLACE : %s", collectionItem[idLabel]))
			newCollectionList = append(newCollectionList, resultMap)
			notInDB = false
		} else {
			println(fmt.Sprintf("OLD V   : %s", collectionItem[idLabel]))
			newCollectionList = append(newCollectionList, collectionItem)
		}
	}

	if notInDB {
		println(fmt.Sprintf("NOT IN DB : %s", idValue))
		newCollectionList = append(newCollectionList, resultMap)
	}

	m.Cache[collection] = newCollectionList
	return nil
}
