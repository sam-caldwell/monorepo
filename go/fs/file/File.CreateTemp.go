package file

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateTemp - Create a temporary file (e.g. /tmp/<name>)
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
	name := filepath.Join(tmpDir)
	if err := fp.valid(&name); err != nil {
		return err
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	(*fp).handle, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	return nil
}
