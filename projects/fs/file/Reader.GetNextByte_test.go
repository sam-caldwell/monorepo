package file

import (
	"testing"
)

func TestReader_GetNextByte(t *testing.T) {
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
		var thisByte byte
		if thisByte, err = reader.GetNextByte(); err != nil {
			t.Fatalf("GetNextByte() returned error: %v", err)
		}
		if thisByte != []byte(testData)[i] {
			t.Fatalf("first byte mismatch at %d (%v)", i, thisByte)
		}
	}

}
