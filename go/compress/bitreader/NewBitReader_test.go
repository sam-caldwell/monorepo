package bitreader

import (
	"bytes"
	"testing"
)

// TestNewBitReader tests the NewBitReader function.
func TestNewBitReader(t *testing.T) {
	// Define a byte slice for testing.
	data := []byte{0b10101010, 0b11001100}

	// Test case 1: Input reader that already implements io.ByteReader.
	mockByteReader := bytes.NewReader(data)
	bitReader := NewBitReader(mockByteReader)
	if _, ok := bitReader.reader.(*bytes.Reader); !ok {
		t.Fatalf("expected *bytes.Reader, got %T", bitReader.reader)
	}

	// Test case 2: Input reader that does not implement io.ByteReader.
	mockReader := bytes.NewBuffer(data)
	bitReader = NewBitReader(mockReader)
	if bitReader == nil {
		t.Fatalf("expected *bitReader, got nil")
	}
	if bitReader.reader == nil {
		t.Fatalf("expected *bitReader.reader, got nil")
	}
	if bitReader.err != nil {
		t.Fatalf("expected *bitReader.err to be nil, got %v", bitReader.err)
	}
}
