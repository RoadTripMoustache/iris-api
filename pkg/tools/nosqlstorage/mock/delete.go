package mock

import "fmt"

// Delete - Delete an element in the selected collection
func (m *MockClient) Delete(collection string, idValue interface{}, idParamLabel string) error {
	var newCache []map[string]interface{}
	for _, element := range m.Cache[collection] {
		if fmt.Sprintf("%v", element[idParamLabel]) != fmt.Sprintf("%v", idValue) {
			newCache = append(newCache, element)
		}
	}

	m.Cache[collection] = newCache
	return nil
}
