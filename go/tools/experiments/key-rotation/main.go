package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/sets/ordered"
	"os"
)

type EncryptionManager struct {
	writableKey  *rsa.PrivateKey
	readOnlyKeys ordered.Set[*rsa.PublicKey]
}

// NewEncryptionManager creates a new EncryptionManager with an initial key pair.
func NewEncryptionManager(keySize uint) (mgr *EncryptionManager, err error) {
	var privateKey *rsa.PrivateKey
	if keySize < 4096 {
		return mgr, fmt.Errorf("key size must be at least 4096")
	}
	if privateKey, err = rsa.GenerateKey(rand.Reader, int(keySize)); err != nil {
		return mgr, err
	}
	mgr = &EncryptionManager{
		writableKey:  privateKey,
		readOnlyKeys: make([]*rsa.PublicKey, 0),
	}
	return mgr, err
}

// RotateKey generates a new key pair, saves the old public key as read-only, and writes the new keys to files.
func (e *EncryptionManager) RotateKey() error {
	// Move the current writable public key to the read-only list
	publicKey := &e.writableKey.PublicKey
	e.readOnlyKeys = append(e.readOnlyKeys, publicKey)

	// Generate new writable key
	newPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Write the new keys to files
	err = writeKeyToFiles(newPrivateKey)
	if err != nil {
		return err
	}

	// Update the writable key
	e.writableKey = newPrivateKey

	return nil
}

// writeKeyToFiles writes the given private key and corresponding public key to separate PEM files.
func writeKeyToFiles(privateKey *rsa.PrivateKey) error {
	// Write the private key to a file
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return err
	}
	defer func() { _ = privateKeyFile.Close() }()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	if err := pem.Encode(privateKeyFile, privateKeyPem); err != nil {
		return err
	}

	// Write the public key to a file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return err
	}
	defer func() { _ = publicKeyFile.Close() }()

	if err := pem.Encode(publicKeyFile, publicKeyPem); err != nil {
		return err
	}

	return nil
}

func main() {
	var err error
	var manager *EncryptionManager
	if manager, err = NewEncryptionManager(); err != nil {
		ansi.Red().Printf("Error creating EncryptionManager: %v\n", err).Fatal(exit.GeneralError).Reset()
	}
	// Rotate the key (this can be done multiple times as needed)
	if err = manager.RotateKey(); err != nil {
		ansi.Red().Printf("Error rotating key: %v\n", err).Fatal(exit.GeneralError).Reset()
	}

	fmt.Println("Key rotation successful")
}
