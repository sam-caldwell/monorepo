package pkicore

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestCreateCertificate(t *testing.T) {
	// Read CA private key and certificate files
	caPrivateKey, caCertificate, err := CreateCa("US", "Texas", "Sonora", "Asymmetric Effort",
		"PKI", "*.asymmetric-effort.com", "scaldwell@asymmetric-effort.com")
	if err != nil {
		t.Fatalf("Error reading CA: %v", err)
	}

	country := "US"
	state := "Texas"
	city := "ExampleCity"
	org := "ExampleOrg"
	ou := "ExampleOU"
	cn := "example.com"

	t.Run("Create non-CA certificate", func(t *testing.T) {
		var privateKey []byte
		var certificate []byte

		privateKey, certificate, err = CreateCertificate(caPrivateKey, caCertificate, country, state, city, org,
			ou, cn, false)
		if err != nil {
			t.Fatalf("Error creating certificate: %v", err)
		}

		// Decode the PEM blocks to verify the generated private key and certificate
		block, _ := pem.Decode(privateKey)
		if block == nil || block.Type != "EC PRIVATE KEY" {
			t.Fatalf("Invalid private key PEM block")
		}

		block, _ = pem.Decode(certificate)
		if block == nil || block.Type != "CERTIFICATE" {
			t.Fatalf("Invalid certificate PEM block")
		}

		// Parse the generated certificate for additional checks if needed
		parsedCert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			t.Fatalf("Error parsing generated certificate: %v", err)
		}
		if parsedCert.IsCA {
			t.Fatalf("certificate should not be CA")
		}
		if parsedCert.Subject.CommonName != cn {
			t.Fatalf("certificate common name mismatch")
		}
	})
}
