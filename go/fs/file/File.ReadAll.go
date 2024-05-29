package file

// ReadAll - Read the entire content of the file and return the same.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (fp *File) ReadAll() (data []byte, err error) {
	_, err = fp.handle.Read(data)
	return data, err
}
