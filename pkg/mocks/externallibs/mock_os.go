// Package externallibs - Contains all the mocks of external libs.
package externallibs

import (
	"github.com/stretchr/testify/mock"
)

// MockOS - Mocking os
type MockOS struct {
	mock.Mock
}

func (m *MockOS) Exit(code int) {
	m.Called(code)
}
