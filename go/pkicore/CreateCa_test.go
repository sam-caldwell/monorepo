package pkicore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestCreateCa(t *testing.T) {
	country := "US"
	state := "CA"
	city := "San Francisco"
	org := "Example Corp"
	ou := "IT"
	cn := "example.com"
	email := "admin@example.com"

	caPrivateKey, caCertificate, err := CreateCa(country, state, city, org, ou, cn, email)

	if err != nil {
		t.Errorf("Error creating CA: %v", err)
	}

	// Decode private key and certificate from PEM
	block, _ := pem.Decode(caPrivateKey)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		t.Error("Failed to decode private key")
	}

	decodedPrivateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("Failed to parse private key: %v", err)
	}

	block, _ = pem.Decode(caCertificate)
	if block == nil || block.Type != "CERTIFICATE" {
		t.Error("Failed to decode certificate")
	}

	decodedCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Errorf("Failed to parse certificate: %v", err)
	}

	// Perform assertions
	if decodedPrivateKey.PublicKey.X.Cmp(decodedCert.PublicKey.(*ecdsa.PublicKey).X) != 0 ||
		decodedPrivateKey.PublicKey.Y.Cmp(decodedCert.PublicKey.(*ecdsa.PublicKey).Y) != 0 {
		t.Error("Private key and certificate public key mismatch")
	}

	if !decodedCert.IsCA {
		t.Error("Certificate is not marked as CA")
	}

	// Add more assertions based on your requirements
}
