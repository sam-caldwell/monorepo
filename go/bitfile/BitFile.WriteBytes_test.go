package bitfile

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

	const testFile = "/tmp/BitFile.WriteBytes.test.txt"

	originalData := []byte("test data")

	defer func() { _ = os.Remove(testFile) }()

	var err error

	var target BitFile

	if target.file, err = os.Open(testFile); err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}

	if err = target.WriteBytes(originalData); err != nil {
		t.Fatalf("")
	}

}
