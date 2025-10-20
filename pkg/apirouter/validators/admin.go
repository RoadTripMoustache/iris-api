package validators

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/apirouter/utils"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/enum"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/errors"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/services/admin"
)

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
