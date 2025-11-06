package admin

import (
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

var (
	adminMap = map[string]interface{}{
		"user_email": "titiU",
	}

	admin = Admin{
		UserEmail: "titiU",
	}

	adminModel = models.Admin{
		UserEmail: "titiU",
	}
)
