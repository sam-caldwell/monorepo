package file

import "os"

func CreateTempFileWithContent(content []byte) (handle *os.File, err error) {
	handle, err = CreateTempFile()
	defer func() {
		_ = handle.Close()
	}()
	_, err = handle.Write(content)
	return handle, err
}
