package services

import (
	"github.com/RoadTripMoustache/iris_api/pkg/tools/auth"
	"github.com/stretchr/testify/mock"
)

// MockAuthClient - Mocking Auth client
type MockAuthClient struct {
	mock.Mock
}

func (m *MockAuthClient) GetInstance() auth.AuthClient {
	args := m.Called()
	return args.Get(0).(auth.AuthClient)
}
