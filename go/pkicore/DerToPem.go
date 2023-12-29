package pkicore

import (
	"encoding/pem"
	"errors"
)

// DERtoPEM converts a DER-encoded certificate to PEM-encoded certificate.
func DERtoPEM(derCertificate []byte) ([]byte, error) {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derCertificate,
	}
	pemCertificate := pem.EncodeToMemory(block)
	if pemCertificate == nil {
		return nil, errors.New("failed to encode DER to PEM")
	}
	return pemCertificate, nil
}
