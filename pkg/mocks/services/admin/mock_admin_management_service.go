// Package admin contains all the mocked admin services
package admin

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	dbmodels "github.com/RoadTripMoustache/iris_api/pkg/dbmodels/admin"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/stretchr/testify/mock"
)

// MockAdminManagementService - Mocking AdminManagementService
type MockAdminManagementService struct {
	mock.Mock
}

func (m *MockAdminManagementService) GetAdmin(ctx utils.Context, userEmail string) (*dbmodels.Admin, *errors.EnhancedError) {
	args := m.Called(ctx, userEmail)
	var activity *dbmodels.Admin
	if args.Get(0) != nil {
		activity = args.Get(0).(*dbmodels.Admin)
	}
	var enhancedError *errors.EnhancedError
	if args.Get(1) != nil {
		enhancedError = args.Get(1).(*errors.EnhancedError)
	}
	return activity, enhancedError
}

func (m *MockAdminManagementService) GetAdmins(ctx utils.Context) ([]*dbmodels.Admin, *errors.EnhancedError) {
	args := m.Called(ctx)
	var admins []*dbmodels.Admin
	if args.Get(0) != nil {
		admins = args.Get(0).([]*dbmodels.Admin)
	}
	var enhancedError *errors.EnhancedError
	if args.Get(1) != nil {
		enhancedError = args.Get(1).(*errors.EnhancedError)
	}
	return admins, enhancedError
}

func (m *MockAdminManagementService) AddAdmin(ctx utils.Context, userEmail string) *errors.EnhancedError {
	args := m.Called(ctx, userEmail)
	var enhancedError *errors.EnhancedError
	if args.Get(0) != nil {
		enhancedError = args.Get(0).(*errors.EnhancedError)
	}
	return enhancedError
}

func (m *MockAdminManagementService) DeleteAdmin(ctx utils.Context, userEmail string) *errors.EnhancedError {
	args := m.Called(ctx, userEmail)
	var enhancedError *errors.EnhancedError
	if args.Get(0) != nil {
		enhancedError = args.Get(0).(*errors.EnhancedError)
	}
	return enhancedError
}
