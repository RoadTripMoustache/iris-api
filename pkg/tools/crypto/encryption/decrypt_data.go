package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
)

func (e *EncryptionClient) DecryptData(data string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("base64 decode error: %w", err)
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, e.PrivateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("decryption error: %w", err)
	}
	return string(plaintext), nil
}
