package services

import (
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

// MockNoSQLStorageInstance - Mocking NoSQLStorageInstance
type MockNoSQLStorageInstance struct {
	mock.Mock
}

func (m *MockNoSQLStorageInstance) GetDocumentsOrderBy(collection string, orderBy bson.D, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	args := m.Called(collection, orderBy, limit, offset, filters)
	return args.Get(0).([]map[string]interface{})
}

func (m *MockNoSQLStorageInstance) GetDocuments(collection string, limit *int, offset *int, filters []utils.Filter) []map[string]interface{} {
	args := m.Called(collection, limit, offset, filters)
	return args.Get(0).([]map[string]interface{})
}

func (m *MockNoSQLStorageInstance) GetFirstDocument(collection string, filters []utils.Filter) []map[string]interface{} {
	args := m.Called(collection, filters)
	return args.Get(0).([]map[string]interface{})
}

func (m *MockNoSQLStorageInstance) Count(collection string, filters []utils.Filter) *int {
	args := m.Called(collection, filters)
	return args.Get(0).(*int)
}

func (m *MockNoSQLStorageInstance) Add(collection string, data interface{}) error {
	args := m.Called(collection, data)
	if args.Get(0) == nil {
		var emptyError error
		return emptyError
	}
	return args.Get(0).(error)
}

func (m *MockNoSQLStorageInstance) Delete(collection string, idValue interface{}, idParamLabel string) error {
	args := m.Called(collection, idValue, idParamLabel)
	if args.Get(0) == nil {
		var emptyError error
		return emptyError
	}
	return args.Get(0).(error)
}

func (m *MockNoSQLStorageInstance) Update(collection string, idValue interface{}, idLabel string, data interface{}) error {
	args := m.Called(collection, idValue, idLabel, data)
	if args.Get(0) == nil {
		var emptyError error
		return emptyError
	}
	return args.Get(0).(error)
}

func (m *MockNoSQLStorageInstance) GetRandomDocuments(collection string, sampleSize int) []map[string]interface{} {
	args := m.Called(collection, sampleSize)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]map[string]interface{})
}

func (m *MockNoSQLStorageInstance) GetRandomDocumentsWithFilter(collection string, sampleSize int, filters map[string]interface{}) []map[string]interface{} {
	args := m.Called(collection, sampleSize, filters)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]map[string]interface{})
}
