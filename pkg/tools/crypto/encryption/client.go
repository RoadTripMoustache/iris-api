// Package encryption contains all the code related to encryption.
package encryption

import (
	"crypto/rsa"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/config"
	"github.com/RoadTripMoustache/guide_nestor_api/pkg/tools/logging"
)

type EncryptionClient struct {
	PrivateKey *rsa.PrivateKey
}

func New() *EncryptionClient {
	key, err := rsaConfigSetup(
		*config.GetConfigs().Database.Mongo.PrivateKey,
		*config.GetConfigs().Database.Mongo.PublicKey,
	)

	if err != nil {
		logging.Error(err, nil)
		return nil
	}

	return &EncryptionClient{
		key,
	}
}
