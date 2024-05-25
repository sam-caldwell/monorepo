package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"io"
)

// GetSize - Return file size (in bytes)
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//		 (c) 2023 Sam Caldwell.  MIT License
func (fp *File) GetSize() (sz int64, err error) {
	if fp.handle == nil {
		return 0, fmt.Errorf(errors.NotInitialized)
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	pos, err := fp.handle.Seek(0, SeekFromCurrent)
	if err != nil {
		return 0, err
	}
	sz, err = fp.handle.Seek(0, SeekFromEnd)
	if err != nil {
		return 0, err
	}
	_, err = fp.handle.Seek(pos, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return sz, err
}
