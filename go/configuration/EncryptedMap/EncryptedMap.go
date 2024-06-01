package configuration

import (
	"sync"
)

// EncryptedMap - A class for handling encrypted configuration files.
type EncryptedMap struct {
	//lock - locking mechanism
	lock sync.RWMutex
	//gpgDir - directory where keys reside
	gpgDir string
	//identity - PGP encryption key id
	identity string
	//data - ciphertext
	data []byte
	//fileName - string identifying config file
	fileName string
}
