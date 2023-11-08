package file

/*
 * BitFile.WriteUint32() tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file contains the unit test(s) for BitFile.WriteUint32()
 */
import (
	"bytes"
	"os"
	"testing"
)

func TestWriteUint32(t *testing.T) {
	var bitFile BitFile
	var tempFileName string
	func() {
		// Create a temporary file for testing
		tempFile, err := os.CreateTemp("", "bitfile32_test")
		if err != nil {
			t.Fatalf("Error creating temp file: %v", err)
		}
		tempFileName = tempFile.Name()
		defer func() { _ = tempFile.Close() }()
	}()
	// Close the BitFile to ensure the data is flushed to the file
	defer bitFile.Close()
	defer func() {
		_ = os.Remove(tempFileName)
	}()
	if err := bitFile.OpenRead(&tempFileName); err != nil {
		t.Fatal(err)
	}
	// Write a uint32 to the BitFile
	err := bitFile.WriteUint32(123456)
	if err != nil {
		t.Fatalf("Error writing uint32 to BitFile: %v", err)
	}

	func() {
		// Read the content of the temporary file
		actual, err := os.ReadFile(tempFileName)
		if err != nil {
			t.Fatalf("Error reading temp file: %v", err)
		}
		// Check if the content of the file matches the expected value
		expected := []byte{64, 226, 1, 0} // Little-endian representation of 123456
		if !bytes.Equal(actual[0:len(expected)], expected) {
			t.Errorf("Unexpected content in temp file. Expected: %v, Got: %v",
				expected, actual)
		}
	}()
}
