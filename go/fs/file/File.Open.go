package file

import (
	"fmt"
	"os"
)

// Open - Open the given file and return its file handle.
//
//		    The goal is to create an abstraction layer so changes to golang will not require
//		    significant effort to update all projects.
//
//	     Throw an error if--
//	         The file is already open
//	         The filename is invalid.
//	         There is an error opening the file.
//
//	     (c) 2023 Sam Caldwell.  MIT License
func (fp *File) Open(name string, flags int, perm os.FileMode) (err error) {
	if fp.handle != nil {
		return fmt.Errorf("file is already open")
	}
	if err := fp.valid(&name); err != nil {
		return err
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	//example flag: os.O_RDWR|os.O_CREATE|os.O_TRUNC
	(*fp).handle, err = os.OpenFile(name, flags, perm)
	return nil
}
