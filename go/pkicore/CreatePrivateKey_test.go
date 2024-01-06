package pkicore

import (
	"crypto/rsa"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {
	tests := []struct {
		name     string
		keySize  int
		expected func(*rsa.PrivateKey, error)
	}{
		{
			name:    "ValidKeySize",
			keySize: 2048,
			expected: func(privateKey *rsa.PrivateKey, err error) {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if privateKey == nil {
					t.Error("Private key is nil")
				}
			},
		},
		{
			name:    "InvalidKeySize",
			keySize: 0, // An invalid key size
			expected: func(privateKey *rsa.PrivateKey, err error) {
				if err == nil {
					t.Error("Expected error, but got nil")
				}
				if privateKey != nil {
					t.Error("Private key should be nil")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privateKey, err := CreatePrivateKey(tt.keySize)
			tt.expected(privateKey, err)
		})
	}
}
