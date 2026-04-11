package gpgCrypto

import (
	"fmt"
	"golang.org/x/crypto/openpgp"
	"os"
	"path/filepath"
	"strings"
)

func LoadPublicKeys(keyDir string, keyIdentity string) ([]*openpgp.Entity, error) {
	var publicKeys []*openpgp.Entity

	// Open the keyring directory
	dirEntries, err := os.ReadDir(keyDir)
	if err != nil {
		return nil, fmt.Errorf("reading key directory: %w", err)
	}

	// Iterate over the files in the directory
	for _, file := range dirEntries {
		if file.IsDir() {
			continue
		}

		keyringPath := filepath.Join(keyDir, " private-keys-v1.d", file.Name())

		if !strings.HasSuffix(keyringPath, ".key") {
			continue
		}

		// Open the keyring file
		keyringFile, err := os.Open(keyringPath)
		if err != nil {
			return nil, fmt.Errorf("opening keyring file(%s): %w", keyringPath, err)
		}
		// Read the keyring
		keyring, err := openpgp.ReadKeyRing(keyringFile)
		if err != nil {
			return nil, fmt.Errorf("reading keyring file(%s): %w", keyringPath, err)
		}
		_ = keyringFile.Close()

		// Find keys with matching identity
		for _, entity := range keyring {
			for _, identity := range entity.Identities {
				if strings.Contains(identity.Name, keyIdentity) {
					publicKeys = append(publicKeys, entity)
				}
			}
		}
	}

	return publicKeys, nil
}
