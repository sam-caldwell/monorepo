package huffman

import (
	"reflect"
	"testing"
)

func TestGenerateHuffmanCodes(t *testing.T) {
	tests := []struct {
		name     string
		tree     *node
		expected map[byte]string
	}{
		{
			name: "Empty",
			tree: &node{},
			expected: map[byte]string{
				0: "0",
			},
		}, {
			name: "Single character",
			tree: &node{Character: 'a', Frequency: 1},
			expected: map[byte]string{
				'a': "0",
			},
		}, {
			name: "Two characters",
			tree: &node{
				Frequency: 2,
				Left: &node{
					Character: 'a',
					Frequency: 1,
				},
				Right: &node{
					Character: 'b',
					Frequency: 1,
				},
			},
			expected: map[byte]string{
				'a': "0",
				'b': "1",
			},
		}, {
			name: "Three characters with different frequencies",
			tree: &node{
				Frequency: 5,
				Left: &node{
					Character: 'a',
					Frequency: 2,
				},
				Right: &node{
					Frequency: 3,
					Left: &node{
						Character: 'b',
						Frequency: 1,
					},
					Right: &node{
						Character: 'c',
						Frequency: 2,
					},
				},
			},
			expected: map[byte]string{
				'a': "0",
				'b': "10",
				'c': "11",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes := make(map[byte]string)
			generateHuffmanCodes(tt.tree, "", codes)
			if !reflect.DeepEqual(codes, tt.expected) {
				t.Errorf("generateHuffmanCodes() = %v, expected %v", codes, tt.expected)
			}
		})
	}
}
