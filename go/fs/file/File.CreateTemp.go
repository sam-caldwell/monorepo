package file

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

// CreateTemp - Create a temporary file (e.g. /tmp/<name>) using uuid.
//
//		 The goal is to create an abstraction layer so changes to golang will not require
//		 significant effort to update all projects.
//
//	     File will be open in write mode.
//
//	     Throw an error if--
//		         The file is already open
//		         The filename is invalid.
//		         There is an error opening the file.
//
//		 (c) 2023 Sam Caldwell.  MIT License
func (fp *File) CreateTemp(perm os.FileMode) (err error) {
	if fp.handle != nil {
		return fmt.Errorf("file is already open")
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	name := filepath.Join(tmpDir, uuid.New().String())
	if err := fp.valid(&name); err != nil {
		return err
	}
	(*fp).handle, err = os.OpenFile(name, FlagReadWrite|FlagCreate|FlagTruncate, perm)
	return nil
}
