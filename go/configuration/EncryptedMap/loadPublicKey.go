package configuration

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"os"
	"path/filepath"
)

// loadPublicKey - load the public gpg key.
//
//	(c) 2023 Sam Caldwell.  MIT License
func loadPublicKey(keysDir, identity string) (key *crypto.Key, err error) {
	var files []os.DirEntry
	if files, err = os.ReadDir(keysDir); err != nil {
		return nil, fmt.Errorf(errFailedToReadKeysDir, err)
	}
	for _, file := range files {
		if !file.IsDir() {
			var (
				keyPath string
				keyData []byte
			)
			keyPath = filepath.Join(keysDir, file.Name())
			if keyData, err = os.ReadFile(keyPath); err != nil {
				return nil, fmt.Errorf(errFailedToReadKeyFile, keyPath, err)
			}
			if key, err = crypto.NewKeyFromArmored(string(keyData)); err != nil {
				continue // Not a valid key, try the next file
			}
			// Check if this key matches the identity
			entity := key.GetEntity()

			for _, ident := range entity.Identities {
				if ident.UserId.Email == identity || ident.UserId.Name == identity {
					return key, nil // Found the matching key
				}
			}
		}
	}
	return nil, fmt.Errorf(errNoKeyNotFound, identity)
}
