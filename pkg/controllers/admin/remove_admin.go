// Package admin contains all the functions exposed by the Admin controller.
package admin

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/enum"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	"github.com/RoadTripMoustache/iris_api/pkg/models/admin"
	adminsvc "github.com/RoadTripMoustache/iris_api/pkg/services/admin"
)

// RemoveAdmin - Remove an user from the admin list
func RemoveAdmin(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	payload := admin.Admin{}
	if err := utils.BodyFormatter(ctx.Body, &payload); err != nil {
		return nil, errors.New(enum.BadRequest, err)
	}

	adm, _ := adminsvc.GetAdmin(ctx, payload.UserEmail)
	if adm == nil {
		return utils.PrepareResponse(nil)
	}

	e := adminsvc.DeleteAdmin(ctx, payload.UserEmail)
	if e != nil {
		return nil, e
	}

	return utils.PrepareResponse(nil)
}
