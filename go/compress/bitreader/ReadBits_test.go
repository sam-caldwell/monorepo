package bitreader

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"testing"
)

func TestBitReader_ReadBits(t *testing.T) {
	var (
		data       []byte
		mockReader *mockByteReader
		br         BitReader
	)
	t.Run("Setup test data and reader", func(t *testing.T) {
		// Create a MockByteReader with test data.
		data = []byte{0b10101010, 0b11001100}
		mockReader = &mockByteReader{data: data}
		br.reader = mockReader
	})

	t.Run("Test reading 3 bits", func(t *testing.T) {
		// Test reading bits.
		bits := br.ReadBits(3)
		if err := br.Error(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if expectedBits := uint64(0b101); bits != expectedBits {
			t.Fatalf("expected %b, got %b", expectedBits, bits)
		}
	})

	t.Run("Test reading 5 bits", func(t *testing.T) {
		// Test reading more bits.
		bits := br.ReadBits(5)
		if err := br.Error(); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if expectedBits := uint64(0b01010); bits != expectedBits {
			t.Fatalf("expected %b, got %b", expectedBits, bits)
		}
	})

	t.Run("Test reading too many bits (expect error)", func(t *testing.T) {
		expectedError := fmt.Errorf(errors.BoundsCheckError)
		_ = br.ReadBits(65)
		err := br.Error()
		if err == nil {
			t.Fatalf("expected error not found")
		}
		if err.Error() != expectedError.Error() {
			t.Fatalf("expected error: %v, got: %v", expectedError, err)
		}
	})

	t.Run("Test reading past the end", func(t *testing.T) {
		expectedError := fmt.Errorf("end of data")
		_ = br.ReadBits(60)
		err := br.Error()
		if err == nil {
			t.Fatalf("expected error not found")
		}
		if err.Error() != expectedError.Error() {
			t.Fatalf("expected error: %v, got: %v", expectedError, err)
		}
	})
}
