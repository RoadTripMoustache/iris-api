// Package crypto contains all the code related to cryptography.
package crypto

import (
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/crypto/encryption"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/crypto/mock"
)

var encryptionInstance CryptoClient = nil

// GetInstance - Get the instance of the datasource client.
func GetInstance() CryptoClient {
	if encryptionInstance == nil {
		if config.GetConfigs().Database.Mock.Enabled {
			encryptionInstance = mock.New()
		} else {
			encryptionInstance = encryption.New()
		}
	}
	return encryptionInstance
}
