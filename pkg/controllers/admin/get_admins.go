// Package admin contains all the functions exposed by the Admin controller.
package admin

import (
	apiUtils "github.com/RoadTripMoustache/iris_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/controllers/utils"
	"github.com/RoadTripMoustache/iris_api/pkg/errors"
	adminsvc "github.com/RoadTripMoustache/iris_api/pkg/services/admin"
)

// GetAdmins returns the list of admins
func GetAdmins(ctx apiUtils.Context) ([]byte, *errors.EnhancedError) {
	admList, e := adminsvc.GetAdmins(ctx)
	if e != nil {
		return nil, e
	}
	return utils.PrepareResponse(admList)
}
