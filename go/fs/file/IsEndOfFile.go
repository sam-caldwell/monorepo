package file

import (
	"errors"
	"io"
)

// isEndOfFile - evaluate an error and return whether it signals end of file
func isEndOfFile(err error) (bool, error) {
	if errors.Is(err, io.EOF) {
		return true, nil
	}
	return false, err
}
