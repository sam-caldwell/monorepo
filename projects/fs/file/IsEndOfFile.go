package file

/*
 * IsEndOfFile()
 * (c) 2023 Sam Caldwell. See License.txt
 *
 * Safely return boolean true if EOF is reached.
 */

import (
	"errors"
	"io"
)

// IsEndOfFile - evaluate an error and return whether it signals end of file
func IsEndOfFile(err error) (bool, error) {
	if errors.Is(err, io.EOF) {
		return true, nil
	}
	return false, err
}
