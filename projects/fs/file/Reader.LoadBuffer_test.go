package file

import (
	"testing"
)

func TestReader_LoadBuffer_Noop(t *testing.T) {
	const (
		bufferSize   = 4
		testFileName = "test_file.txt"
		testData     = "this is a test data string"
	)
	var err error
	var reader Reader

	file := createTempFile(t, testFileName, []byte(testData))
	defer deleteTempFile(t, file)

	// Test the basic state after an initial load
	if _, err = reader.Open(file.Name(), bufferSize); err != nil {
		t.Fatalf("Error opening file: %v", err)
	}
	if reader.bufferPosition != 0 {
		t.Fail()
	}
	if reader.filePosition == 0 {
		t.Fatal("filePosition should be advanced past 0")
	}
	if reader.filePosition < uint64(len(reader.buffer)) {
		t.Fatal("filePosition should reflect a single buffer-length from 0")
	}

	reader.bufferPosition = len(reader.buffer)

	if err = reader.LoadBuffer(len(reader.buffer)); err != nil {
		t.Fatalf("error loading buffer (2): %v", err)
	}
	if reader.bufferPosition != 0 {
		t.Fail()
	}
	if reader.filePosition != uint64(2*len(reader.buffer)) {
		t.Fatalf("file position should be 2x buffer since we have read buffer twice.\n"+
			"Actual  : %d\n"+
			"Expected: %d\n"+
			"Buffer  : '%s'", reader.filePosition, 2*len(reader.buffer), string(reader.buffer))
	}
}
