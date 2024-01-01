package pkicore

import (
	"crypto/elliptic"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {
	// Call the function to create a private key
	privateKey, err := CreatePrivateKey()

	// Check for any errors during key generation
	if err != nil {
		t.Fatalf("Error generating private key: %v", err)
	}

	// Verify that the private key uses the P-256 elliptic curve
	expectedCurve := elliptic.P256()
	if privateKey.Curve != expectedCurve {
		t.Fatalf("Expected curve %v, got %v", expectedCurve, privateKey.Curve)
	}
}
