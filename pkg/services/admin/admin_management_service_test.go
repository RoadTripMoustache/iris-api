package admin

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	dbmodels "github.com/RoadTripMoustache/iris_api/pkg/dbmodels/admin"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/mocks/services"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ----- GetAdmin ----- //
func Test_GetAdmin(t *testing.T) {
	testCases := []struct {
		caseDesc       string
		dbResult       []map[string]interface{}
		queryParams    map[string][]string
		expectedResult *dbmodels.Admin
		exprectedErr   *errors.EnhancedError
	}{
		{
			caseDesc:       "Empty DB",
			dbResult:       []map[string]interface{}{},
			queryParams:    nil,
			expectedResult: nil,
			exprectedErr:   nil,
		},
		{
			caseDesc:    "Not empty DB",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_id": "toto",
				},
			},
			expectedResult: &dbmodels.Admin{
				UserID: "toto",
			},
			exprectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseDesc, func(t *testing.T) {
			// --- Given
			userID := "toto_uuid"
			ctx := apiUtils.Context{
				UserID:      userID,
				QueryParams: testCase.queryParams,
			}

			mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
			mNoSQLStorageInstance.
				On("GetFirstDocument", dbmodels.AdminCollectionName,
					[]utils.Filter{{
						Param:    dbmodels.AdminUserIDLabel,
						Operator: "==",
						Value:    "totoU",
					}}).
				Return(testCase.dbResult)

			mockNoSQLStorageService := new(services.MockNoSQLStorage)
			mockNoSQLStorageService.On("GetInstance").Return(mNoSQLStorageInstance)
			noSQLStorageGetInstance = mockNoSQLStorageService.GetInstance

			// --- When
			result, err := GetAdmin(ctx, "totoU")

			// --- Then
			mNoSQLStorageInstance.AssertExpectations(t)
			mockNoSQLStorageService.AssertExpectations(t)
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
	}{
		{
			caseDesc:       "Empty DB",
			dbResult:       []map[string]interface{}{},
			queryParams:    nil,
			expectedResult: nil,
			exprectedErr:   nil,
		},
		{
			caseDesc:    "1 item",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_id": "toto",
				},
			},
			expectedResult: []*dbmodels.Admin{
				{
					UserID: "toto",
				},
			},
			exprectedErr: nil,
		},
		{
			caseDesc:    "2 items",
			queryParams: nil,
			dbResult: []map[string]interface{}{
				{
					"user_id": "toto",
				},
				{
					"user_id": "titi",
				},
			},
			expectedResult: []*dbmodels.Admin{
				{
					UserID: "toto",
				},
				{
					UserID: "titi",
				},
			},
			exprectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.caseDesc, func(t *testing.T) {
			// --- Given
			userID := "toto_uuid"
			pageNumber := 1
			pageOffset := 0
			ctx := apiUtils.Context{
				UserID:      userID,
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
		UserID: "toto",
	}
	expectedResult := dbmodels.Admin{
		UserID: "titi",
	}

	mNoSQLStorageInstance := new(services.MockNoSQLStorageInstance)
	var nilError error
	mNoSQLStorageInstance.
		On("Add", dbmodels.AdminCollectionName, expectedResult.ToMap()).
		Return(nilError)

	mNoSQLStorageInstance.
		On("GetFirstDocument", dbmodels.AdminCollectionName,
			[]utils.Filter{{
				Param:    dbmodels.AdminUserIDLabel,
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
		On("Delete", dbmodels.AdminCollectionName, "titi", dbmodels.AdminUserIDLabel).
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
