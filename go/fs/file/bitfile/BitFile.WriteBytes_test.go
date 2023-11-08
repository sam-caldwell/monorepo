package file

/*
 * CRSCE BitFile.WriteBytes() Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the BitFile.WriteBytes() method which performs a bitwise buffered write to disk.
 */

import (
	"os"
	"testing"
)

func TestBitFile_WriteBytes(t *testing.T) {

	const testFileName = "/tmp/BitFile.WriteBytes.test.txt"

	var err error
	var target BitFile

	originalData := []byte("test data")

	target.file, err = os.Create(testFileName)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() {
		_ = target.file.Close()
		_ = os.Remove(testFileName)
	}()

	if err = target.WriteBytes(originalData); err != nil {
		t.Fatalf("error writing bytes.")
	}

}
