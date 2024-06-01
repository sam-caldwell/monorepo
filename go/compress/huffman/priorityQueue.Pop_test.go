package huffman

import (
    "testing"
)

func TestPriorityQueue_Pop(t *testing.T) {
	tests := []struct {
		name     string
		pq       priorityQueue
		expected *node
		afterPop priorityQueue
	}{
		{
			name: "Pop from queue with multiple elements",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
				{Character: 'c', Frequency: 3},
			},
			expected: &node{
				Character: 'c',
				Frequency: 3,
			},
			afterPop: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
		},
		{
			name: "Pop from queue with one element",
			pq: priorityQueue{
				{Character: 'a', Frequency: 1},
			},
			expected: &node{
				Character: 'a',
				Frequency: 1,
			},
			afterPop: priorityQueue{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thisNode := tt.pq.Pop().(*node)
			if thisNode != nil && (thisNode.Character != tt.expected.Character || thisNode.Frequency != tt.expected.Frequency) {
				t.Errorf("Pop() = %v, expected %v", thisNode, tt.expected)
			}
			if len(tt.pq) != len(tt.afterPop) {
				t.Fatalf("After Pop(), length of queue = %d, expected %d", len(tt.pq), len(tt.afterPop))
			}
			for i := range tt.pq {
				if tt.pq[i].Character != tt.afterPop[i].Character || tt.pq[i].Frequency != tt.afterPop[i].Frequency {
					t.Errorf("After Pop(), pq[%d] = %v, expected %v", i, tt.pq[i], tt.afterPop[i])
				}
			}
		})
	}
}
