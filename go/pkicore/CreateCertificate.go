package pkicore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
)

func CreateCertificate(caPrivate, caCert []byte, country, state, city, org, ou, cn string,
	isCa bool) (privateKey, certificate []byte, err error) {

	var privKey *ecdsa.PrivateKey
	var privKeyBytes []byte

	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate private key: %v", err)
	}
	privKeyBytes, err = x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encode private key: %v", err)
	}
	privateKey = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privKeyBytes})
	subject := pkix.Name{
		Country:            []string{country},
		Province:           []string{state},
		Locality:           []string{city},
		Organization:       []string{org},
		OrganizationalUnit: []string{ou},
		CommonName:         cn,
	}
	_, err = x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject:            subject,
		SignatureAlgorithm: x509.ECDSAWithSHA256,
	}, privKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create CSR: %v", err)
	}
	caPrivateDer, _ := pem.Decode(caPrivate)
	caPrivKey, err := x509.ParseECPrivateKey(caPrivateDer.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse CA private key: %v", err)
	}
	caCertPem, _ := pem.Decode(caCert)
	if caCertPem == nil {
		return nil, nil, fmt.Errorf("failed to decode CA certificate")
	}
	caCertParsed, err := x509.ParseCertificate(caCertPem.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse CA certificate: %v", err)
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &x509.Certificate{
		Subject:               subject,
		NotBefore:             caCertParsed.NotBefore,
		NotAfter:              caCertParsed.NotAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		SerialNumber:          big.NewInt(2),
		BasicConstraintsValid: true,
		IsCA:                  isCa,
	}, caCertParsed, &privKey.PublicKey, caPrivKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create signed certificate: %v", err)
	}
	certificate = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	return privateKey, certificate, nil
}
