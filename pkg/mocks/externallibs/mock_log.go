// Package externallibs - Contains all the mocks of external libs.
package externallibs

import (
	"github.com/stretchr/testify/mock"
)

// MockLog - Mocking Log
type MockLog struct {
	mock.Mock
}

func (m *MockLog) Error(err error, args map[string]interface{}) {
	m.Called(err, args)
}

func (m *MockLog) Fatal(err error, args map[string]interface{}) {
	m.Called(err, args)
}

func (m *MockLog) Info(msg string, args map[string]interface{}) {
	m.Called(msg, args)
}

func (m *MockLog) Warn(msg string, args map[string]interface{}) {
	m.Called(msg, args)
}
