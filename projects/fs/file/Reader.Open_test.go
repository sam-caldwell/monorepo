package file

import (
	"testing"
)

func TestReader_Open(t *testing.T) {
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
	if reader.buffer == nil {
		t.Fatalf("buffer is not initialized.")
	}
	if len(reader.buffer) != bufferSize {
		t.Fatalf("Buffer size is wrong")
	}
	if string(reader.buffer[0:4]) != testData[0:4] {
		t.Fatalf("buffer did not load")
	}
	if reader.bufferPosition != 0 {
		t.Fail()
	}
	if reader.bufferBitPos != 0 {
		t.Fail()
	}
	if reader.filePosition == 0 {
		t.Fatal("filePosition should be advanced past 0")
	}
	if reader.filePosition < uint64(len(reader.buffer)) {
		t.Fatal("filePosition should reflect a single buffer-length from 0")
	}
}
