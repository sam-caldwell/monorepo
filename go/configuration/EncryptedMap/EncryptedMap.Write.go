package configuration

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

// Write - Write the current state to disk.
//
//	The internal state (data) is ciphertext and is written as-is
//	to disk.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (em *EncryptedMap) Write() (err error) {
	if em.data == nil {
		return fmt.Errorf(errors.EmptySet)
	}
	em.lock.Lock()
	defer em.lock.Unlock()
	if strings.TrimSpace(em.fileName) == words.EmptyString {
		return fmt.Errorf(errors.NotInitialized)
	}
	var configFile file.File
	defer configFile.Close()
	if err = configFile.Open(em.fileName, wflags, file.OwnerRead|file.OwnerWrite); err != nil {
		return err
	}
	_, err = configFile.WriteBytes(em.data)
	return err
}
