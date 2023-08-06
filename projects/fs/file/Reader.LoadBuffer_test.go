package file

import (
	"testing"
)

func TestReader_LoadBuffer_Noop(t *testing.T) {
	//const (
	//	bufferSize   = 4
	//	testFileName = "test_file.txt"
	//	testData     = "this is a test data string"
	//)
	//var err error
	//
	//file := createTempFile(t, testFileName, []byte(testData))
	//defer deleteTempFile(t, file)
	//
	//// Test Full Buffer load
	//var r Reader
	//if _, err = r.Open(file.Name(), bufferSize); err != nil {
	//	t.Fatalf("Failed to open file: %v", err)
	//}
	//defer r.Close()
	//
	//// Read the first 4 bytes from the file into the buffer
	//if err = r.LoadBuffer(); err != nil {
	//	t.Fatalf("Failed to load buffer: %v", err)
	//}
}
