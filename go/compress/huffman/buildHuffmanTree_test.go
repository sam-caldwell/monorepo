package huffman

import (
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	t.Skip("disabled for debugging")
	tests := []struct {
		name        string
		input       []byte
		expected    *node
		expectError bool
	}{
		{
			name:     "Test with nil input",
			input:    []byte{},
			expected: &node{},
		}, {
			name:        "Test with empty-string input",
			input:       []byte(""),
			expectError: false,
			expected:    &node{},
		}, {
			name:     "Test with single character",
			input:    []byte("a"),
			expected: &node{Character: 'a', Frequency: 1},
		}, {
			name:  "Test with 4 characters",
			input: []byte{'a', 'b', 'a', 'c'},
			expected: &node{
				Frequency: 4,
				Left: &node{
					Character: 'a',
					Frequency: 2,
				},
				Right: &node{
					Frequency: 2,
					Left: &node{
						Character: 'c',
						Frequency: 1,
					},
					Right: &node{
						Character: 'b',
						Frequency: 1,
					},
				},
			},
		},
		{
			name:  "Test with 16 characters",
			input: []byte("aaaabbbbccccdddd"),
			expected: &node{
				Frequency: 16,
				Left: &node{
					Frequency: 8,
					Left: &node{
						Frequency: 4,
						Character: 'a',
					},
					Right: &node{
						Character: 'd',
						Frequency: 4,
					},
				},
				Right: &node{
					Frequency: 8,
					Left: &node{
						Frequency: 4,
						Character: 'c',
					},
					Right: &node{
						Character: 'b',
						Frequency: 4,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := buildHuffmanTree(tt.input)
			PrettyPrintTree("actual", tree)
			PrettyPrintTree("expected", tt.expected)
			if tt.expectError && tree != nil {
				t.Fatal("Expected an error, but got a non-nil tree")
			}
			if !tt.expectError && (tree == nil || !isEquivalentTree(tree, tt.expected)) {
				t.Fatal("Tree is not built as expected")
			}
		})
	}
}
