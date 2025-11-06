package admin

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/config"
	dbmodels "github.com/RoadTripMoustache/iris_api/pkg/dbmodels/admin"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/mocks/services"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

var NoAdminConfig = config.Config{
	Admin: config.AdminConfig{
		DefaultList: []string{},
	},
}

var WithAdminConfig = config.Config{
	Admin: config.AdminConfig{
		DefaultList: []string{
			"toto@toto.com",
		},
	},
}

// ----- GetAdmin ----- //
func Test_GetAdmin(t *testing.T) {
	testCases := []struct {
		caseDesc       string
		dbResult       []map[string]interface{}
		queryParams    map[string][]string
		expectedResult *dbmodels.Admin
		exprectedErr   *errors.EnhancedError
		mockedConfig   config.Config
	}{
		{
			caseDesc:       "Empty DB and no default admin",
			dbResult:       []map[string]interface{}{},
			queryParams:    nil,
			expectedResult: nil,
			exprectedErr:   nil,
			mockedConfig:   NoAdminConfig,
		},
		{
			caseDesc:    "Not empty DB and no default admin",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_email": "toto@toto.com",
				},
			},
			expectedResult: &dbmodels.Admin{
				UserEmail: "toto@toto.com",
			},
			exprectedErr: nil,
			mockedConfig: NoAdminConfig,
		},
		{
			caseDesc:    "Empty DB and with default admin",
			dbResult:    []map[string]interface{}{},
			queryParams: nil,
			expectedResult: &dbmodels.Admin{
				UserEmail: "toto@toto.com",
			},
			exprectedErr: nil,
			mockedConfig: WithAdminConfig,
		},
		{
			caseDesc:    "Not empty DB and with default admin",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_email": "toto@toto.com",
				},
			},
			expectedResult: &dbmodels.Admin{
				UserEmail: "toto@toto.com",
			},
			exprectedErr: nil,
			mockedConfig: WithAdminConfig,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseDesc, func(t *testing.T) {
			// --- Given
			userEmail := "toto@toto.com"
			ctx := apiUtils.Context{
				UserEmail:   userEmail,
				QueryParams: testCase.queryParams,
			}

			configGetConfigs = func() config.Config {
				return testCase.mockedConfig
			}

			mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
			mNoSQLStorageInstance.
				On("GetFirstDocument", dbmodels.AdminCollectionName,
					[]utils.Filter{{
						Param:    dbmodels.AdminUserEmailLabel,
						Operator: "==",
						Value:    "toto@toto.com",
					}}).
				Return(testCase.dbResult)

			mockNoSQLStorageService := new(services.MockNoSQLStorage)
			mockNoSQLStorageService.On("GetInstance").Return(mNoSQLStorageInstance)
			noSQLStorageGetInstance = mockNoSQLStorageService.GetInstance

			// --- When
			result, err := GetAdmin(ctx, "toto@toto.com")

			// --- Then
			if len(testCase.mockedConfig.Admin.DefaultList) == 0 {
				mNoSQLStorageInstance.AssertExpectations(t)
				mockNoSQLStorageService.AssertExpectations(t)
			}
			assert.Equal(t, testCase.exprectedErr, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

// ----- GetAdmins ----- //
func Test_GetAdmins(t *testing.T) {
	testCases := []struct {
		caseDesc       string
		dbResult       []map[string]interface{}
		queryParams    map[string][]string
		expectedResult []*dbmodels.Admin
		exprectedErr   *errors.EnhancedError
		mockedConfig   config.Config
	}{
		{
			caseDesc:       "Empty DB and no default admin",
			dbResult:       []map[string]interface{}{},
			queryParams:    nil,
			expectedResult: nil,
			exprectedErr:   nil,
			mockedConfig:   NoAdminConfig,
		},
		{
			caseDesc:    "1 item in DB",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_email": "toto",
				},
			},
			expectedResult: []*dbmodels.Admin{
				{
					UserEmail: "toto",
				},
			},
			exprectedErr: nil,
			mockedConfig: NoAdminConfig,
		},
		{
			caseDesc:    "2 items in DB",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_email": "toto",
				},
				{
					"user_email": "titi",
				},
			},
			expectedResult: []*dbmodels.Admin{
				{
					UserEmail: "toto",
				},
				{
					UserEmail: "titi",
				},
			},
			exprectedErr: nil,
			mockedConfig: NoAdminConfig,
		},
		{
			caseDesc:    "1 item in default admin",
			queryParams: nil,
			dbResult:    nil,
			expectedResult: []*dbmodels.Admin{
				{
					UserEmail: "toto@toto.com",
				},
			},
			exprectedErr: nil,
			mockedConfig: WithAdminConfig,
		},
		{
			caseDesc:    "2 items in DB and default admin",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_email": "toto",
				},
				{
					"user_email": "titi",
				},
			},
			expectedResult: []*dbmodels.Admin{
				{
					UserEmail: "toto",
				},
				{
					UserEmail: "titi",
				},
				{
					UserEmail: "toto@toto.com",
				},
			},
			exprectedErr: nil,
			mockedConfig: WithAdminConfig,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseDesc, func(t *testing.T) {
			// --- Given
			userEmail := "toto@toto.com"
			pageNumber := 1
			pageOffset := 0
			ctx := apiUtils.Context{
				UserEmail:   userEmail,
				QueryParams: testCase.queryParams,
				Pagination: apiUtils.Pagination{
					PageNumber: &pageNumber,
					PageSize:   &pageNumber,
				},
			}
			var filters []utils.Filter

			mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
			mNoSQLStorageInstance.
				On("GetDocuments", dbmodels.AdminCollectionName, &pageNumber, &pageOffset, filters).
				Return(testCase.dbResult)

			mockNoSQLStorageService := new(services.MockNoSQLStorage)
			mockNoSQLStorageService.On("GetInstance").Return(mNoSQLStorageInstance)
			noSQLStorageGetInstance = mockNoSQLStorageService.GetInstance

			configGetConfigs = func() config.Config {
				return testCase.mockedConfig
			}

			// --- When
			result, err := GetAdmins(ctx)

			// --- Then
			mNoSQLStorageInstance.AssertExpectations(t)
			mockNoSQLStorageService.AssertExpectations(t)
			assert.Equal(t, testCase.exprectedErr, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

// ----- AddAdmin ----- //
func Test_AddAdmin(t *testing.T) {
	// --- Given
	ctx := apiUtils.Context{
		UserEmail: "toto",
	}
	expectedResult := dbmodels.Admin{
		UserEmail: "titi",
	}

	mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
	var nilError error
	mNoSQLStorageInstance.
		On("Add", dbmodels.AdminCollectionName, expectedResult.ToMap()).
		Return(nilError)

	mNoSQLStorageInstance.
		On("GetFirstDocument", dbmodels.AdminCollectionName,
			[]utils.Filter{{
				Param:    dbmodels.AdminUserEmailLabel,
				Operator: "==",
				Value:    "titi",
			}}).
		Return([]map[string]interface{}{})

	mockNoSQLStorageService := new(services.MockNoSQLStorage)
	mockNoSQLStorageService.On("GetInstance").Return(mNoSQLStorageInstance)
	noSQLStorageGetInstance = mockNoSQLStorageService.GetInstance

	// --- When
	err := AddAdmin(ctx, "titi")

	// --- Then
	mockNoSQLStorageService.AssertExpectations(t)
	mNoSQLStorageInstance.AssertExpectations(t)
	assert.Nil(t, err)
}

// ----- DeleteAdmin ----- //
func Test_DeleteAdmin(t *testing.T) {
	// --- Given
	ctx := apiUtils.Context{}

	mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
	var nilError error
	mNoSQLStorageInstance.
		On("Delete", dbmodels.AdminCollectionName, "titi", dbmodels.AdminUserEmailLabel).
		Return(nilError)

	mockNoSQLStorageService := new(services.MockNoSQLStorage)
	mockNoSQLStorageService.On("GetInstance").Return(mNoSQLStorageInstance)
	noSQLStorageGetInstance = mockNoSQLStorageService.GetInstance

	// --- When
	result := DeleteAdmin(ctx, "titi")

	// --- Then
	mockNoSQLStorageService.AssertExpectations(t)
	mNoSQLStorageInstance.AssertExpectations(t)
	assert.Nil(t, result)
}
