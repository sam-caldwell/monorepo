package huffmanPrint

import (
	"testing"
)

func TestGetBranch(t *testing.T) {
	tests := []struct {
		name     string
		isTail   bool
		expected string
	}{
		{
			name:     "Test with isTail true",
			isTail:   true,
			expected: "└",
		},
		{
			name:     "Test with isTail false",
			isTail:   false,
			expected: "├",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getBranch(tt.isTail)
			if result != tt.expected {
				t.Errorf("getBranch(%v) = %v, expected %v", tt.isTail, result, tt.expected)
			}
		})
	}
}
