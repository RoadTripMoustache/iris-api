// Package admin contains all the DB models for admin resources
package admin

import (
	"encoding/json"
	models "github.com/RoadTripMoustache/iris_api/pkg/models/admin"
)

const (
	AdminCollectionName string = "admins"
	AdminUserEmailLabel string = "user_email"
)

type Admin struct {
	UserEmail string `json:"user_email"`
}

func (a *Admin) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_email": a.UserEmail,
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
		UserEmail: admin.UserEmail,
	}

	return modelAdmin
}
