// Package validators contains all the validation tools for the router
package validators

import (
	"github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/services/admin"
)

// IsAdmin - Validate if the current user is an admin. Returns an EnhancedError if not.
func IsAdmin(ctx utils.Context) *errors.EnhancedError {
	user, err := admin.GetAdmin(ctx, ctx.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New(enum.AuthUnauthorized, nil)
	}
	return nil
}
