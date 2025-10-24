package admin

import (
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

var (
	adminMap = map[string]interface{}{
		"user_id": "titiU",
	}

	admin = Admin{
		UserID: "titiU",
	}

	adminModel = models.Admin{
		UserID: "titiU",
	}
)
