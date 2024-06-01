package huffman

import "testing"

func TestPriorityQueue_Less(t *testing.T) {
	tests := []struct {
		name     string
		pq       priorityQueue
		i        int
		j        int
		expected bool
	}{
		{
			name: "First node has lower frequency",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
			i:        0,
			j:        1,
			expected: true,
		},
		{
			name: "First node has higher frequency",
			pq: priorityQueue{
				{Character: 'a', Frequency: 3},
				{Character: 'b', Frequency: 2},
			},
			i:        0,
			j:        1,
			expected: false,
		},
		{
			name: "Both nodes have equal frequency",
			pq: priorityQueue{
				{Character: 'a', Frequency: 2},
				{Character: 'b', Frequency: 2},
			},
			i:        0,
			j:        1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pq.Less(tt.i, tt.j)
			if result != tt.expected {
				t.Errorf("priorityQueue.Less() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
