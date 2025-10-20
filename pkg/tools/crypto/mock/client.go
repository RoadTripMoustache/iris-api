// Package mock contains the mocked crypto tool
package mock

type MockClient struct {
}

func New() *MockClient {
	return &MockClient{}
}
