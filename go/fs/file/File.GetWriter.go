package file

import "io"

// GetWriter - Return the file handle as io.Writer
//
//	(c) 2023 Sam Caldwell.  MIT License
func (fp *File) GetWriter() io.Writer {
	return fp.handle
}
