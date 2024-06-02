package bitreader

import (
	"fmt"
	"testing"
)

func TestBitReader_ReadBit(t *testing.T) {
	var (
		data           []byte
		mockReader     *mockByteReader
		br             BitReader
		expectedValues []bool
	)
	t.Run("Setup test data and reader", func(t *testing.T) {
		data = []byte{0b10101010, 0b11001100}
		expectedValues = []bool{
			//00     01    02     03    04     05    06     07    08    09     10     11    12    13     14     15
			true, false, true, false, true, false, true, false, true, true, false, false, true, true, false, false,
		}
		mockReader = &mockByteReader{data: data}
		br.reader = mockReader
	})

	t.Run("Test reading bits", func(t *testing.T) {
		for pos, expected := range expectedValues {
			bit := br.ReadBit()
			if err := br.Error(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if bit != expected {
				t.Fatalf("expected bit to be %v, got %v (pos %d)", expected, bit, pos)
			}
		}
	})

	t.Run("Test reading past the end", func(t *testing.T) {
		expectedError := fmt.Errorf("end of data")
		for i := 0; i < 20; i++ {
			_ = br.ReadBit()
			err := br.Error()
			if err == nil {
				t.Fatalf("expected error not found")
			}
			if err.Error() != expectedError.Error() {
				t.Fatalf("expected error: %v, got: %v", expectedError, err)
			}
		}
	})
}
