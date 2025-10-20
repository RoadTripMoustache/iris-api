package mock

import "encoding/json"

// Add - Add data in the selected collection
func (m *MockClient) Add(collection string, data interface{}) error {
	if m.Cache[collection] == nil {
		m.Cache[collection] = make([]map[string]interface{}, 0)
	}

	jsonBytes, _ := json.Marshal(data)
	var resultMap map[string]interface{}
	json.Unmarshal(jsonBytes, &resultMap)

	m.Cache[collection] = append(m.Cache[collection], resultMap)
	return nil
}
