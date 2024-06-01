package huffman

import (
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	// t.Skip("disabled for debugging")

	tests := []struct {
		name        string
		input       []byte
		expected    *node
		expectError bool
	}{
		{
			name:     "Test with nil input",
			input:    nil,
			expected: nil,
		}, {
			name:     "Test with empty-string input",
			input:    []byte(""),
			expected: nil,
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
						Character: 'b',
						Frequency: 1,
					},
					Right: &node{
						Character: 'c',
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
						Character: 'a',
						Frequency: 4,
					},
					Right: &node{
						Character: 'b',
						Frequency: 4,
					},
				},
				Right: &node{
					Frequency: 8,
					Left: &node{
						Character: 'c',
						Frequency: 4,
					},
					Right: &node{
						Character: 'd',
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
			if eqTree := isEquivalentTree(tree, tt.expected); eqTree {
				if tt.expectError {
					t.Fatalf("tree is equivalent but error is expected but not found.")
				} else {
					return
				}
			}
		})
	}
}
