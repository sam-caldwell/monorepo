package file

import (
	"testing"
)

func TestReader_GetNextBit(t *testing.T) {
	const (
		bufferSize   = 4
		testFileName = "test_file.txt"
		testData     = "this is a test data string"
	)
	var err error
	var reader Reader
	defer reader.Close()

	file := createTempFile(t, testFileName, []byte(testData))
	defer deleteTempFile(t, file)

	if _, err = reader.Open(file.Name(), bufferSize); err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	// Verify first buffer full of stuff.
	if string(reader.buffer) != string(testData[0:len(reader.buffer)]) {
		t.Fatalf("buffer slice does not match (1):%s", string(reader.buffer))
	}
	for i := 0; i < len(reader.buffer); i++ {
		var actualByte = []byte(testData)[i] //This is what we expect (byte)
		for bit := 0; bit < 8; bit++ {
			var thisBit bool
			var actualBit = (actualByte>>uint(7-bit))&1 != 0x00 // This is what we expect (bit)

			if thisBit, err = reader.GetNextBit(); err != nil {
				t.Fatalf("GetNextBit() returned error at (%d,%d): %v", i, bit, err)
			}
			if thisBit != actualBit {
				t.Fatalf("first byte mismatch at (%d:%d) (%v)", i, bit, thisBit)
			}
		}
	}

}
