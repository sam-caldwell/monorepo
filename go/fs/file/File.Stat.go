package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"os"
)

// Stat - Return file info
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//		 (c) 2023 Sam Caldwell.  MIT License
func (fp *File) Stat() (os.FileInfo, error) {
	if fp.handle == nil {
		return nil, fmt.Errorf(errors.NotInitialized)
	}
	return fp.handle.Stat()
}
