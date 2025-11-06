// Package admin contains all the functions exposed by the Admin controller.
package admin

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	adminsvc "github.com/RoadTripMoustache/iris_api/pkg/services/admin"
)

// IsCurrentUserAdmin returns whether the authenticated user is an admin.
func IsCurrentUserAdmin(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	if ctx.UserEmail == "" {
		return nil, errors.New(enum.AuthUnauthorized, nil)
	}
	adm, e := adminsvc.GetAdmin(ctx, ctx.UserEmail)
	if e != nil {
		return nil, e
	}
	resp := map[string]bool{"is_admin": adm != nil}
	return utils.PrepareResponse(resp)
}
