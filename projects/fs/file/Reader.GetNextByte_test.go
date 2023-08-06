package file

import (
	"testing"
)

func TestReader_GetNextByte(t *testing.T) {
	//bufferSize := 2
	//expected := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99}
	//
	//// Create a temporary file for testing
	//file, err := os.CreateTemp("", "test_file")
	//if err != nil {
	//	t.Fatalf("Failed to create temporary file: %v", err)
	//}
	//defer func() {
	//	if err = os.Remove(file.Name()); err != nil {
	//		t.Fatalf("Failed to remove file: %v", err)
	//	}
	//}()
	//
	//// Write some data to the temporary file
	//data := expected
	//_, err = file.Write(data)
	//if err != nil {
	//	t.Fatalf("Failed to write data to the temporary file: %v", err)
	//}
	//
	//// Test GetNextByte method
	//var r Reader
	//_, err = r.Open(file.Name(), bufferSize) // Use small buffer size for testing
	//if err != nil {
	//	t.Fatalf("Failed to open file: %v", err)
	//}
	//defer r.Close()
	//
	//// Read the bytes from the file using GetNextByte()
	//for i, expected := range data {
	//	b, err := r.GetNextByte()
	//	if err != nil {
	//		t.Fatalf("Error while reading byte: %v", err)
	//	}
	//	if b != expected {
	//		t.Fatalf("At %d Expected byte 0x%x, but got 0x%x", i, expected, b)
	//	}
	//}
	//
	//// Ensure GetNextByte returns an error when the end of file is reached
	//_, err = r.GetNextByte()
	//if err == nil {
	//	t.Fatal("Expected an error, but got nil")
	//}
}
