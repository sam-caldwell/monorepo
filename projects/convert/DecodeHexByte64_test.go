package convert

import (
	"reflect"
	"testing"
)

func TestDecodeHexByte64(t *testing.T) {
	tests := []struct {
		input       []byte
		expected    [8]byte
		shouldError bool
	}{
		{
			input:       []byte("FF00A5CDFF00A5CD"),
			expected:    [8]byte{255, 0, 165, 205, 255, 0, 165, 205},
			shouldError: false,
		},
		{
			input:       []byte("00FF00FF00FF00FF"),
			expected:    [8]byte{0, 255, 0, 255, 0, 255, 0, 255},
			shouldError: false,
		},
		{
			input:       []byte("GARBAGEINPUT"),
			expected:    [8]byte{},
			shouldError: true,
		},
	}

	for _, test := range tests {
		result, ok := DecodeHexByte64(test.input)

		if test.shouldError {
			if ok {
				t.Errorf("Expected failure for input '%s', but got ok", test.input)
			}
		} else {
			if !ok {
				t.Errorf("Expected ok but failed for input '%s'", test.input)
			} else if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Unexpected result for input '%s'. Expected: %v, Got: %v",
					test.input, test.expected, result)
			}
		}
	}
}
