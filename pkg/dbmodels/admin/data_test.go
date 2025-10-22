package admin

import (
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

var (
	adminMap = map[string]interface{}{
		"user_id": "titiU",
	}

	admin = Admin{
		UserId: "titiU",
	}

	adminModel = models.Admin{
		UserId: "titiU",
	}
)
