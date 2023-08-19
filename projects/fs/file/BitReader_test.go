package file

import (
	"testing"
)

func TestBitReaderInitialState(t *testing.T) {
	// Create a new BitReader instance
	reader := BitReader{}

	// Check initial state properties
	if reader.buffer != nil {
		t.Error("Expected non-nil buffer, but got nil")
	}

	if reader.count != 0 {
		t.Errorf("Expected count to be 0, but got %d", reader.count)
	}

	if reader.readBuffer != nil {
		t.Error("Expected readBuffer to be nil, but got non-nil")
	}

	if reader.done {
		t.Error("Expected done to be false, but got true")
	}
}
