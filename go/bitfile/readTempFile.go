package bitfile

/*
 * readTempFile()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * read the content of a given temp file for bitfile tests.
 */

import "os"

// Helper function to read the content of a file
func readTempFile(file *os.File) ([]byte, error) {
	_, err := file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	content := make([]byte, bufferSize)
	_, err = file.Read(content)
	if err != nil {
		return nil, err
	}

	return content, nil
}
