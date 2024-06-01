package huffman

import "testing"

func TestPriorityQueue_Swap(t *testing.T) {
	tests := []struct {
		name     string
		pq       priorityQueue
		i        int
		j        int
		expected priorityQueue
	}{
		{
			name: "Swap first and second elements",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
			i: 0,
			j: 1,
			expected: priorityQueue{
				{Character: 'b', Frequency: 2},
				{Character: 'a', Frequency: 1},
			},
		},
		{
			name: "Swap same element",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
			i: 0,
			j: 0,
			expected: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
		},
		{
			name: "Swap second and third elements",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
				{Character: 'c', Frequency: 3},
			},
			i: 1,
			j: 2,
			expected: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'c', Frequency: 3},
				{Character: 'b', Frequency: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.i, tt.j)
			for k := range tt.pq {
				if tt.pq[k].Character != tt.expected[k].Character || tt.pq[k].Frequency != tt.expected[k].Frequency {
					t.Errorf("After Swap(), pq[%d] = %v, expected %v", k, tt.pq[k], tt.expected[k])
				}
			}
		})
	}
}
