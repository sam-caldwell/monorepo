package file

import (
	"testing"
)

func TestReader_Close(t *testing.T) {
	// Create a temporary file for testing
	//file, err := os.CreateTemp("", "test_file")
	//if err != nil {
	//	t.Fatalf("Failed to create temporary file: %v", err)
	//}
	//
	//defer func() {
	//	err = os.Remove(file.Name())
	//	if err != nil {
	//		t.Fatalf("Failed to remove file: %v", err)
	//	}
	//}()
	//
	//// Write some data to the temporary file
	//data := []byte("Hello, World!")
	//_, err = file.Write(data)
	//if err != nil {
	//	t.Fatalf("Failed to write data to the temporary file: %v", err)
	//}
	//
	//// Define the buffer size for testing
	//bufferSize := 8
	//
	//// Test Close method
	//r := &Reader{}
	//_, err = r.Open(file.Name(), bufferSize)
	//if err != nil {
	//	t.Fatalf("Failed to open file: %v", err)
	//}
	//
	//// Read a few bytes from the file using GetNextByte()
	//for i := 0; i < 5; i++ {
	//	_, err := r.GetNextByte()
	//	if err != nil {
	//		t.Fatalf("Error while reading byte: %v", err)
	//	}
	//}

	// Call Close method to close the file and reset position indexes
	//reader.Close()

	//// Verify that the file handle is nil after closing
	//if reader.h != nil {
	//	t.Fatal("File handle is not nil after closing")
	//}
	//
	//// Verify that the position index properties (`fPos` and `bitPos`) are reset to 0 after closing
	//if reader.filePosition != 0 {
	//	t.Fatalf("fPos is %d, expected 0", r.filePosition)
	//}
	//if reader.bufferBitPos != 0 {
	//	t.Fatalf("bitPos is %d, expected 0", r.bufferBitPos)
	//}
}
