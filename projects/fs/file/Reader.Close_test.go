package file

import (
	"testing"
)

func TestReader_Close(t *testing.T) {
	const (
		bufferSize   = 4
		testFileName = "test_file.txt"
		testData     = "this is a test data string"
	)
	var err error
	var reader Reader

	file := createTempFile(t, testFileName, []byte(testData))
	defer deleteTempFile(t, file)

	if _, err = reader.Open(file.Name(), bufferSize); err != nil {
		t.Fatalf("Error opening file: %v", err)
	}

	reader.Close()

	if reader.h != nil {
		t.Fatal("File handle is not nil after closing")
	}
}
