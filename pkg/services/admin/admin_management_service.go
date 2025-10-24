// Package admin contains all the admin management service methods.
package admin

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	dbmodels "github.com/RoadTripMoustache/iris_api/pkg/dbmodels/admin"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage"
	nosqlutils "github.com/RoadTripMoustache/iris_api/pkg/tools/nosqlstorage/utils"
)

var (
	noSQLStorageGetInstance = nosqlstorage.GetInstance
)

// GetAdmin retrieves the admin record for a user with the specified ID.
// Parameters:
//   - ctx: The API context containing request information
//   - userID: The ID of the user to check for admin privileges
//
// Returns:
//   - *dbmodels.Admin: The admin record if found, nil otherwise
//   - *errors.EnhancedError: Error information if the operation fails
func GetAdmin(ctx utils.Context, userID string) (*dbmodels.Admin, *errors.EnhancedError) {
	requestFilters := []nosqlutils.Filter{{
		Param:    dbmodels.AdminUserIDLabel,
		Operator: "==",
		Value:    userID,
	}}

	documents := noSQLStorageGetInstance().
		GetFirstDocument(dbmodels.AdminCollectionName, requestFilters)

	var admin *dbmodels.Admin

	for _, doc := range documents {
		admin = dbmodels.AdminFromMap(doc)
	}

	return admin, nil
}

// GetAdmins retrieves all admin records with pagination support.
// Parameters:
//   - ctx: The API context containing request information and pagination settings
//
// Returns:
//   - []*dbmodels.Admin: A slice of admin records
//   - *errors.EnhancedError: Error information if the operation fails
func GetAdmins(ctx utils.Context) ([]*dbmodels.Admin, *errors.EnhancedError) {
	documents := noSQLStorageGetInstance().
		GetDocuments(dbmodels.AdminCollectionName, ctx.Pagination.PageSize, ctx.Pagination.GetOffset(), nil)

	var admins []*dbmodels.Admin

	for _, doc := range documents {
		admin := dbmodels.AdminFromMap(doc)
		admins = append(admins, admin)
	}

	return admins, nil
}

// AddAdmin grants admin privileges to a user with the specified ID.
// If the user already has admin privileges, an error is returned.
// Parameters:
//   - ctx: The API context containing request information
//   - userID: The ID of the user to grant admin privileges to
//
// Returns:
//   - *errors.EnhancedError: Error information if the operation fails or if the user already has admin privileges
func AddAdmin(ctx utils.Context, userID string) *errors.EnhancedError {
	a, eerr := GetAdmin(ctx, userID)
	if eerr != nil {
		return eerr
	}
	if a != nil {
		return errors.New(enum.PermissionsAlreadyGranted, nil)
	}

	newAdmin := dbmodels.Admin{
		UserID: userID,
	}

	err := noSQLStorageGetInstance().Add(dbmodels.AdminCollectionName, newAdmin.ToMap())

	if err != nil {
		logging.Error(err, nil)
		return errors.New(enum.InternalServerError, nil)
	}

	return nil
}

// DeleteAdmin revokes admin privileges from a user with the specified ID.
// Parameters:
//   - ctx: The API context containing request information
//   - userID: The ID of the user to revoke admin privileges from
//
// Returns:
//   - *errors.EnhancedError: Error information if the operation fails
func DeleteAdmin(ctx utils.Context, userID string) *errors.EnhancedError {
	err := noSQLStorageGetInstance().Delete(dbmodels.AdminCollectionName, userID, dbmodels.AdminUserIDLabel)
	if err != nil {
		logging.Error(err, nil)
		return errors.New(enum.InternalServerError, nil)
	}
	return nil
}
