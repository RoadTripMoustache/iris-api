package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

func (e *EncryptionClient) EncryptData(data string) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, &e.PrivateKey.PublicKey, []byte(data))
	if err != nil {
		return "", fmt.Errorf("encryption error: %w", err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
