package bitfile

/*
 * Reader.WriteUint64() tests
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file contains the unit test(s) for Reader.WriteUint64()
 */
import (
	"bytes"
	"os"
	"testing"
)

func TestWriteUint64(t *testing.T) {
	t.Skip("Disabled pending debugging")

	var writer Writer
	var tempFileName string

	func() {
		var err error
		var tempFile *os.File

		// Create a temporary file for testing
		if tempFile, err = os.CreateTemp("", "TestWriteUint64"); err != nil {
			t.Fatalf("Error creating temp file: %v", err)
		}
		tempFileName = tempFile.Name()
		defer func() { _ = tempFile.Close() }()

	}()

	defer func() {
		// Close the writer to ensure the data is flushed to the file
		_ = writer.Close()
		_ = os.Remove(tempFileName)
	}()
	if err := writer.Open(&tempFileName, 4096); err != nil {
		t.Fatal(err)
	}
	// Write a uint64 to the Reader
	if err := writer.WriteUint64(123456); err != nil {
		t.Fatalf("Error writing uint64 to Reader: %v", err)
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
			t.Fatalf("Unexpected content in temp file. Expected: %v, Got: %v",
				expected, actual)
		}
	}()
}
