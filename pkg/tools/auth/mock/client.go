package mock

type MockClient struct {
}

func New() *MockClient {
	return &MockClient{}
}
