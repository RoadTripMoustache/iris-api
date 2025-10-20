// Package auth contains all the code to manage authentication.
package auth

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	firestore "github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/auth/firestore"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/auth/mock"
)

var authInstance AuthClient = nil

// GetInstance - Get the instance of the datasource client.
func GetInstance() AuthClient {
	if authInstance == nil {
		if config.GetConfigs().Firebase.Mock.Enabled {
			authInstance = mock.New()
		} else {
			authInstance = firestore.New()
		}
	}
	return authInstance
}
