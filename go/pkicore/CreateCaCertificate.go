package pkicore

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"time"
)

// GenerateSelfSigned - Create a self-signed certificate CA.
func GenerateSelfSigned(location string, validDays int) (caPrivateKeyPem []byte, caPEM []byte, err error) {

	const (
		caKeySize        = 16384
		rootDomain       = "samcaldwell.net"
		organization     = "pkicore"
		organizationUnit = "certificateAuthorityRoot"
	)

	var caBytes []byte
	var caPrivateKey *rsa.PrivateKey
	var commonName = fmt.Sprintf("%s.%s", uuid.New().String(), rootDomain)

	if validDays < 0 {
		return nil, nil, fmt.Errorf("ttlDays must be greater than 0")
	}

	ca := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			CommonName:         commonName,
			Locality:           []string{location},
			Organization:       []string{organization},
			OrganizationalUnit: []string{organizationUnit},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Duration(validDays) * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caPrivateKey, err = CreatePrivateKey(caKeySize)

	caBytes, err = x509.CreateCertificate(rand.Reader, ca, ca, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		return caPrivateKey, caBytes, err
	}

	caPemBuffer := new(bytes.Buffer)
	err = pem.Encode(caPemBuffer, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		return nil, nil, err
	}
	caPrivKeyPEM := new(bytes.Buffer)
	err = pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey),
	})
	if err != nil {
		return nil, nil, err
	}
	return caPrivKeyPEM.Bytes(), caPemBuffer.Bytes(), err
}
