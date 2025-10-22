package admin

import (
	"encoding/json"
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

const (
	AdminCollectionName string = "admins"
	AdminUserIdLabel    string = "user_id"
)

type Admin struct {
	UserId string `json:"user_id"`
}

func (a *Admin) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id": a.UserId,
	}
}

func AdminFromMap(m map[string]interface{}) *Admin {
	mapByte, _ := json.Marshal(m)
	admin := Admin{}
	json.Unmarshal(mapByte, &admin)

	return &admin
}

func ToAdminModel(admin Admin) models.Admin {
	modelAdmin := models.Admin{
		UserId: admin.UserId,
	}

	return modelAdmin
}
