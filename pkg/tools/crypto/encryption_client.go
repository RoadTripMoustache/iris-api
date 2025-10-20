package crypto

type CryptoClient interface {
	EncryptData(data string) (string, error)

	DecryptData(data string) (string, error)
}
