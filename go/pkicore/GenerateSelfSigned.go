package pkicore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"log"
	"math/big"
	"time"
)

// GenerateSelfSigned - Create a self-signed certificate for use as CA or stand-alone certificate.
func GenerateSelfSigned(commonName string, organization string, organizationUnit string, location string, isCa bool,
	usage x509.KeyUsage, extUsage []x509.ExtKeyUsage, ttlDays int) (caPrivateKey *ecdsa.PrivateKey,
	certificateDer []byte, err error) {

	if ttlDays < 0 {
		return nil, nil, fmt.Errorf("ttlDays must be greater than 0")
	}

	// Generate a self-signed CA certificate and private key
	caPrivateKey, _ = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	// Create a certificate template we will sign later.
	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			CommonName:         commonName,
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationUnit},
			Locality:           []string{location},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, ttlDays),
		KeyUsage:              usage,
		ExtKeyUsage:           extUsage,
		BasicConstraintsValid: true,
		IsCA:                  isCa,
	}

	// Sign the certificate template
	certificateDer, err = x509.CreateCertificate(rand.Reader, caTemplate, caTemplate,
		&caPrivateKey.PublicKey, caPrivateKey)

	if err != nil {
		log.Fatalf("Error creating certificate: %v", err)
	}

	return caPrivateKey, certificateDer, err
}
