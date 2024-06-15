package convert

import (
	"math"
	"testing"
)

func TestInt32ToBytes(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expectErr bool
		expected  []byte
	}{
		{"MinInt32", math.MinInt32, false, []byte{0x80, 0x00, 0x00, 0x00}},
		{"MaxInt32", math.MaxInt32, false, []byte{0x7F, 0xFF, 0xFF, 0xFF}},
		{"Zero", 0, false, []byte{0x00, 0x00, 0x00, 0x00}},
		{"Positive Value", 12345678, false, []byte{0x00, 0xBC, 0x61, 0x4E}},
		{"Negative Value", -12345678, false, []byte{0xFF, 0x43, 0x9E, 0xB2}},
		{"Out of Bounds Positive", math.MaxInt32 + 1, true, nil},
		{"Out of Bounds Negative", math.MinInt32 - 1, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Int32ToBytes(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !tt.expectErr && !equal(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
