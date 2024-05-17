package asymmetricCrypto

import (
	"crypto/rsa"
	"io"
)

// NewReader - Return a
func NewReader(in io.Reader) (out *io.Reader, err error) {
	return out, err
}

// Manager - Manage the keys and encryption lifecycle
type Manager struct {
	privateKey map[string]*rsa.PrivateKey
	publicKey  map[string]*rsa.PublicKey
}

func (mgr *Manager) EncryptionReader() (out io.Reader, err error) {
	return out, err
}
