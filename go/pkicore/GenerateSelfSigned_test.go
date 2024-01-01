package pkicore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"os"
	"testing"
)

func TestGenerateSelfSigned(t *testing.T) {
	const (
		caCertFile       = "test_ca_cert.pem"
		commonName       = "TestCA"
		organization     = "TestOrg"
		organizationUnit = "TestOrgUnit"
		location         = "TestLocation"
	)
	var err error
	var certificateDer []byte
	var caPrivateKey *ecdsa.PrivateKey

	t.Run("Generate self-signed certificate and private key", func(t *testing.T) {
		caPrivateKey, certificateDer, err = GenerateSelfSigned(commonName, organization, organizationUnit, location,
			true, x509.KeyUsageKeyEncipherment|x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign,
			[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, 10)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		if caPrivateKey == nil {
			t.Fatalf("Expected non-nil caPrivateKey, but got nil")
		}
		if (certificateDer == nil) || (len(certificateDer) == 0) {
			t.Fatalf("Expected non-empty certificateDer, but got empty")
		}
	})

	t.Run("Parse the generated certificate for further validation", func(t *testing.T) {
		parsedCert, err := x509.ParseCertificate(certificateDer)
		if err != nil {
			t.Fatalf("Failed to parse certificate: %v", err)
		}
		if parsedCert.Subject.CommonName != commonName {
			t.Fatalf("Expected CommonName %s, but got %s", commonName, parsedCert.Subject.CommonName)
		}
		if len(parsedCert.Subject.OrganizationalUnit) != 1 {
			t.Fatalf("expected one OU record but got %d records", len(parsedCert.Subject.OrganizationalUnit))
		}
		if parsedCert.Subject.OrganizationalUnit[0] != organizationUnit {
			t.Fatalf("Expected OU %s, but got %s", organizationUnit, parsedCert.Subject.OrganizationalUnit)
		}
		if !parsedCert.IsCA {
			t.Fatalf("Expected CA flag")
		}
		if parsedCert.Subject.Locality[0] != location {
			t.Fatalf("Expected location %s but got %s", location, parsedCert.Subject.Locality[0])
		}
		if parsedCert.Subject.Organization[0] != organization {
			t.Fatalf("Expected location %s but got %s", organization, parsedCert.Subject.Organization[0])
		}
		if !parsedCert.BasicConstraintsValid {
			t.Fatalf("Expected BasicConstraintsValid true")
		}
		if parsedCert.KeyUsage != x509.KeyUsageKeyEncipherment|x509.KeyUsageDigitalSignature|x509.KeyUsageCertSign {
			t.Fatalf("keyusage mismatch")
		}
		if parsedCert.ExtKeyUsage[0] != x509.ExtKeyUsageServerAuth &&
			parsedCert.ExtKeyUsage[1] != x509.ExtKeyUsageClientAuth {
			t.Fatalf("ExtKeyUsage[] mismatch")
		}
	})
}

func removeTestFile(filename string) error {
	// Remove the temporary test file
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
