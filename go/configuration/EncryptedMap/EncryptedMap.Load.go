package configuration

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// Load - Load the encrypted file into memory.
//
// The configuration file is loaded from disk and into memory.
// The data is still encrypted.  If no data is read or the file
// does not exist, this method will return an error.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (em *EncryptedMap) Load() (err error) {
	if err = em.loadData(); err != nil {
		return err
	}
	if err = em.loadIdentity(); err != nil {
		return err
	}
	return err
}

func (em *EncryptedMap) loadData() (err error) {
	em.lock.Lock()
	defer em.lock.Unlock()
	if strings.TrimSpace(em.fileName) == words.EmptyString {
		return fmt.Errorf(errors.NotInitialized)
	}
	if !file.Existsp(&em.fileName) {
		return fmt.Errorf(errors.NotFound)
	}
	var configFile file.File
	defer configFile.Close()
	if err = configFile.Open(em.fileName, rflags, perms); err != nil {
		return err
	}
	if em.data, err = configFile.ReadAll(); err != nil {
		return err
	}
	if em.data == nil || strings.TrimSpace(string(em.data)) == words.EmptyString {
		return fmt.Errorf(errors.EmptySet)
	}
	return err
}

func (em *EncryptedMap) loadIdentity() (err error) {
	return err
}
