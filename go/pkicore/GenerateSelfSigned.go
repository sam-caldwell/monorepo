package pkicore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func GenerateSelfSigned(caCertFile, commonName, organization, organizationUnit,
	location string) (caPrivateKey *ecdsa.PrivateKey, certificateDer []byte) {

	var err error
	var fh *os.File

	// Generate a self-signed CA certificate and private key
	caPrivateKey, _ = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			CommonName:         commonName,
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationUnit},
			Locality:           []string{location},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certificateDer, err = x509.CreateCertificate(rand.Reader, caTemplate, caTemplate,
		&caPrivateKey.PublicKey, caPrivateKey)

	if err != nil {
		log.Fatalf("Error creating certificate: %v", err)
	}

	fh, err = os.Create(caCertFile)
	if err != nil {
		log.Fatalf("error saving CA certificate to file: %v", err)
	}

	if err = pem.Encode(fh, &pem.Block{Type: "CERTIFICATE", Bytes: certificateDer}); err != nil {
		log.Fatalf("error encoding cert: %v", err)
	}

	if err = fh.Close(); err != nil {
		log.Fatalf("error closing file: %v", err)
	}

	return caPrivateKey, certificateDer
}
