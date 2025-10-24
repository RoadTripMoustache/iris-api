// Package admin contains all the DB models for admin resources
package admin

import (
	"encoding/json"
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

const (
	AdminCollectionName string = "admins"
	AdminUserIDLabel    string = "user_id"
)

type Admin struct {
	UserID string `json:"user_id"`
}

func (a *Admin) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id": a.UserID,
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
		UserID: admin.UserID,
	}

	return modelAdmin
}
