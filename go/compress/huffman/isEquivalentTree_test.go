package huffman

import (
	"testing"
)

func TestIsEquivalentTree(t *testing.T) {
	tests := []struct {
		name     string
		tree1    *node
		tree2    *node
		expected bool
	}{
		{
			name:     "Both trees are nil",
			tree1:    nil,
			tree2:    nil,
			expected: true,
		},
		{
			name:     "One tree is nil, the other is not",
			tree1:    &node{Character: 'a', Frequency: 1},
			tree2:    nil,
			expected: false,
		},
		{
			name:     "Trees with same structure and values",
			tree1:    &node{Character: 'a', Frequency: 1},
			tree2:    &node{Character: 'a', Frequency: 1},
			expected: true,
		},
		{
			name: "Trees with different structures",
			tree1: &node{
				Character: 'a',
				Frequency: 1,
				Left:      &node{Character: 'b', Frequency: 2},
				Right:     &node{Character: 'c', Frequency: 3},
			},
			tree2: &node{
				Character: 'a',
				Frequency: 1,
				Left:      &node{Character: 'b', Frequency: 2},
			},
			expected: false,
		},
		{
			name: "Trees with same structure but different values",
			tree1: &node{
				Character: 'a',
				Frequency: 1,
				Left:      &node{Character: 'b', Frequency: 2},
				Right:     &node{Character: 'c', Frequency: 3},
			},
			tree2: &node{
				Character: 'a',
				Frequency: 1,
				Left:      &node{Character: 'b', Frequency: 2},
				Right:     &node{Character: 'd', Frequency: 3},
			},
			expected: false,
		},
		{
			name: "Complex equivalent trees",
			tree1: &node{
				Character: 'a',
				Frequency: 6,
				Left: &node{
					Character: 'b',
					Frequency: 3,
					Left:      &node{Character: 'd', Frequency: 1},
					Right:     &node{Character: 'e', Frequency: 2},
				},
				Right: &node{Character: 'c', Frequency: 3},
			},
			tree2: &node{
				Character: 'a',
				Frequency: 6,
				Left: &node{
					Character: 'b',
					Frequency: 3,
					Left:      &node{Character: 'd', Frequency: 1},
					Right:     &node{Character: 'e', Frequency: 2},
				},
				Right: &node{Character: 'c', Frequency: 3},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEquivalentTree(tt.tree1, tt.tree2)
			if result != tt.expected {
				t.Errorf("isEquivalentTree(%v, %v) = %v, expected %v", tt.tree1, tt.tree2, result, tt.expected)
			}
		})
	}
}
