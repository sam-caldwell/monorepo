package pkicore

import (
	"crypto/rand"
	"crypto/rsa"
)

// CreatePrivateKey generates a new elliptic curve private key using P521.
func CreatePrivateKey(keySize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
