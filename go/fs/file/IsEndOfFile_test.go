package file

/*
 * file.IsEndOfFile() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the IsEndOfFile() function
 */

import (
	"errors"
	"io"
	"testing"
)

func TestIsEndOfFile(t *testing.T) {
	func() {
		// Test when error is io.EOF
		err := io.EOF
		isEOF, err := IsEndOfFile(err)

		if !isEOF {
			t.Errorf("Expected IsEndOfFile to return true for io.EOF, but got false")
		}

		if err != nil {
			t.Errorf("Expected error to be nil for io.EOF, but got %v", err)
		}

	}()
	func() {
		// Test when error is not io.EOF
		err := errors.New("some other error")
		isEOF, err := IsEndOfFile(err)

		if isEOF {
			t.Errorf("Expected IsEndOfFile to return false for non-io.EOF error, but got true")
		}

		if err == nil {
			t.Errorf("Expected error to be non-nil for non-io.EOF error, but got %v", err)
		}
	}()
}
