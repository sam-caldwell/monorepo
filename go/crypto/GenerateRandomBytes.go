package crypto

import (
	"crypto/rand"
)

// GenerateRandomBytes generates a random byte slice of the specified length
func GenerateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}
