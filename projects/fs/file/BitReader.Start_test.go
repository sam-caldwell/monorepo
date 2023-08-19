package file

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
	"sync"
	"testing"
)

func TestBitReaderStart(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}

	defer func() { _ = tempFile.Close() }()
	defer func() { _ = os.Remove(tempFile.Name()) }()

	// Write some content to the temporary file
	content := []byte{
		0x00, // 00000000
		0xAA, // 10101010
		0x55, // 01010101
		0xFF, // 11111111
		0x00, // 00000000
	}
	_, err = tempFile.Write(content)
	t.Logf("tempFile: %s", tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to write content to temporary file: %v", err)
	}
	if sz := GetFileSize(tempFile); sz != int64(len(content)) {
		t.Fatalf("GetFileSize() mismatch: %d", sz)
	}

	// Create a BitReader instance
	reader := NewBitReader(1024, 10)

	var wg sync.WaitGroup
	wg.Add(1)

	// Start reading the file with the BitReader
	go reader.Start(tempFile.Name(), &wg)

	// Collect the bits from the buffer and compare with expected
	expectedBits := []bool{
		false, false, false, false, false, false, false, false, // 00000000 (position 0 - 7)
		true, false, true, false, true, false, true, false, //     10101010 (position 8 - 15)
		false, true, false, true, false, true, false, true, //     01010101 (position 16 - 23)
		true, true, true, true, true, true, true, true, //         11111111 (position 24 - 31)
		false, false, false, false, false, false, false, false, // 00000000 (position 32 - 39)
	}

	for position, expectedBit := range expectedBits {
		receivedBit := <-reader.buffer
		if receivedBit != expectedBit {
			t.Errorf("Expected bit %t, but received %t (position %d)", expectedBit, receivedBit, position)
		}
	}
	if !reader.done {
		ansi.Red().Println("reader should be done").Reset()
	}
	// Wait for the reading to complete
	wg.Wait()
	if reader.count != int64(len(expectedBits)) {
		t.Fatalf("count: %d not equal len(expectedBits): %d",
			len(reader.readBuffer), len(expectedBits))
	}

	// Check that the buffer is closed and done is set to true
	select {
	case _, ok := <-reader.buffer:
		if ok {
			t.Error("Expected buffer to be closed")
		}
	default:
		t.Error("Expected buffer to be closed")
	}

	if !reader.done {
		t.Error("Expected done to be true after reading")
	}
}
