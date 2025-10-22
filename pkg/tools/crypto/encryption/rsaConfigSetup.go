package encryption

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/RoadTripMoustache/iris_api/pkg/tools/logging"
	"os"
)

func rsaConfigSetup(rsaPrivateKeyLocation, rsaPublicKeyLocation string) (*rsa.PrivateKey, error) {
	if rsaPrivateKeyLocation == "" {
		logging.Error(fmt.Errorf("no RSA Key given, generating temp one"), nil)
		return nil, fmt.Errorf("nul")
	}

	priv, err := os.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		logging.Error(fmt.Errorf("no RSA private key found, generating temp one"), nil)
		return nil, fmt.Errorf("nul")
	}

	privPem, _ := pem.Decode(priv)
	var privPemBytes []byte
	if privPem.Type != "RSA PRIVATE KEY" {
		logging.Error(fmt.Errorf("RSA private key is of the wrong type :%s", privPem.Type), nil)
	}
	privPemBytes = privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
			logging.Error(err, nil)
			return nil, fmt.Errorf("nul")
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		logging.Error(err, nil)
		return nil, fmt.Errorf("nul")
	}

	pub, err := os.ReadFile(rsaPublicKeyLocation)
	if err != nil {
		logging.Error(fmt.Errorf("no RSA public key found, generating temp one"), nil)
		return nil, fmt.Errorf("nul")
	}
	pubPem, _ := pem.Decode(pub)
	if pubPem == nil {
		logging.Error(fmt.Errorf("use `ssh-keygen -f id_rsa.pub -e -m pem > id_rsa.pem` to generate the pem encoding of your RSA public :rsa public key not in pem format: %s", rsaPublicKeyLocation), nil)

		return nil, fmt.Errorf("nul")
	}
	if pubPem.Type != "PUBLIC KEY" {
		logging.Error(fmt.Errorf("RSA public key is of the wrong type, Pem Type :%s", pubPem.Type), nil)
		return nil, fmt.Errorf("nul")
	}

	if parsedKey, err = x509.ParsePKIXPublicKey(pubPem.Bytes); err != nil {
		logging.Error(err, nil)
		return nil, fmt.Errorf("nul")
	}

	var pubKey *rsa.PublicKey
	if pubKey, ok = parsedKey.(*rsa.PublicKey); !ok {
		logging.Error(err, nil)
		return nil, fmt.Errorf("nul")
	}

	privateKey.PublicKey = *pubKey

	return privateKey, nil
}
