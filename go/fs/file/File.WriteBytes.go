package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// WriteBytes - Write bytes to the file
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) WriteBytes(content []byte) (int, error) {
	if fp.handle == nil {
		return 0, fmt.Errorf(errors.NotInitialized)
	}
	n, err := fp.handle.Write(content)
	return n, err
}
