package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"io"
)

// IsEndOfFile - Return boolean if end of file
//
//	     The goal is to create an abstraction layer so changes to golang will not require
//	     significant effort to update all projects.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) IsEndOfFile() (eof bool, err error) {
	var currentPosition int64
	if fp.handle == nil {
		return true, fmt.Errorf(errors.NotInitialized)
	}
	fp.lock.Lock()
	defer fp.lock.Unlock()
	currentPosition, err = fp.handle.Seek(0, io.SeekCurrent)
	if err != nil {
		return true, err
	}
	fileInfo, err := fp.handle.Stat()
	if err != nil {
		return true, err
	}
	return currentPosition == fileInfo.Size(), nil
}
