package pkicore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

// CreatePrivateKey generates a new elliptic curve private key.
func CreatePrivateKey() (*ecdsa.PrivateKey, error) {
	// Using the P-256 elliptic curve, you can choose other curves like P-384 or P-521
	curve := elliptic.P256()

	// Generate a new private key using the chosen elliptic curve
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	return privateKey, nil
}
