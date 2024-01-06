package pkicore

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"
)

func TestCreateCaCertificate(t *testing.T) {
	location := "TestLocation"
	validDays := uint(365) // One year validity for testing
	caPrivateKeyPem, caPEM, err := CreateCaCertificate(location, validDays)
	if err != nil {
		t.Fatalf("Error creating CA certificate: %v", err)
	}
	// Validate the PEM-encoded CA certificate
	block, _ := pem.Decode(caPEM)
	if block == nil || block.Type != "CERTIFICATE" {
		t.Error("Invalid PEM-encoded certificate")
	}
	// Parse the certificate to check its properties
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("Error parsing PEM-encoded certificate: %v", err)
	}
	// Validate the certificate properties
	if !cert.IsCA {
		t.Error("Certificate is not marked as CA")
	}
	// Validate the validity period
	now := time.Now()
	if now.Before(cert.NotBefore) || now.After(cert.NotAfter) {
		t.Error("Invalid certificate validity period")
	}
	// Validate the PEM-encoded private key
	block, _ = pem.Decode(caPrivateKeyPem)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		t.Error("Invalid PEM-encoded private key")
	}
}
