package file

import (
	"testing"
)

func TestNewBitReader(t *testing.T) {
	// Define the desired buffer sizes
	bitBufferSize := 10
	readBufferSize := 1024

	// Create a new BitReader using the NewBitReader function
	reader := NewBitReader(bitBufferSize, readBufferSize)

	// Check the properties of the created BitReader
	if reader == nil {
		t.Error("Expected non-nil BitReader, but got nil")
	}

	if reader.count != 0 {
		t.Errorf("Expected count to be 0, but got %d", reader.count)
	}

	if reader.done {
		t.Error("Expected done to be false, but got true")
	}

	// Check the size of the buffer channel
	if cap(reader.buffer) != bitBufferSize {
		t.Errorf("Expected buffer capacity to be %d, but got %d", bitBufferSize, cap(reader.buffer))
	}

	// Check the size of the readBuffer slice
	if len(reader.readBuffer) != readBufferSize {
		t.Errorf("Expected readBuffer length to be %d, but got %d", readBufferSize, len(reader.readBuffer))
	}
}
