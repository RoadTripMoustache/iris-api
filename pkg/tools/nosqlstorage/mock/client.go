package mock

type MockClient struct {
	Cache map[string][]map[string]interface{}
}

func New() *MockClient {
	return &MockClient{
		Cache: make(map[string][]map[string]interface{}),
	}
}
