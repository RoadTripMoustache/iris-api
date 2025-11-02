package services

import (
	"firebase.google.com/go/auth"
	"github.com/stretchr/testify/mock"
)

// MockAuthClientInstance - Mocking Auth client instance
type MockAuthClientInstance struct {
	mock.Mock
}

func (m *MockAuthClientInstance) VerifyIDToken(idToken string) *auth.Token {
	args := m.Called(idToken)
	return args.Get(0).(*auth.Token)
}

func (m *MockAuthClientInstance) GetUser(userID string) *auth.UserRecord {
	args := m.Called(userID)
	return args.Get(0).(*auth.UserRecord)
}

func (m *MockAuthClientInstance) GetUserByEmail(email string) *auth.UserRecord {
	args := m.Called(email)
	return args.Get(0).(*auth.UserRecord)
}

func (m *MockAuthClientInstance) DeleteUser(userID string) error {
	args := m.Called(userID)
	if args.Get(0) == nil {
		var emptyError error
		return emptyError
	}
	return args.Get(0).(error)
}
