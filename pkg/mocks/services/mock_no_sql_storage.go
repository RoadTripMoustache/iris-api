package services

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/nosqlstorage"
	"github.com/stretchr/testify/mock"
)

// MockNoSQLStorage - Mocking NoSQLStorage service
type MockNoSQLStorage struct {
	mock.Mock
}

func (m *MockNoSQLStorage) GetInstance() nosqlstorage.NoSQLClient {
	args := m.Called()
	return args.Get(0).(nosqlstorage.NoSQLClient)
}
