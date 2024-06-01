package huffman

import (
	"testing"
)

func TestPriorityQueue_Push(t *testing.T) {
	tests := []struct {
		name      string
		pq        priorityQueue
		element   *node
		afterPush priorityQueue
	}{
		{
			name:      "Push to empty queue",
			pq:        priorityQueue{},
			element:   &node{Character: 'a', Frequency: 1},
			afterPush: priorityQueue{{Character: 'a', Frequency: 1}},
		},
		{
			name: "Push to non-empty queue",
			pq:   priorityQueue{{Character: 'a', Frequency: 1}},
			element: &node{
				Character: 'b',
				Frequency: 2,
			},
			afterPush: priorityQueue{
				{Character: 'a', Frequency: 1},
				{Character: 'b', Frequency: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.element)
			if len(tt.pq) != len(tt.afterPush) {
				t.Fatalf("After Push(), length of queue = %d, expected %d", len(tt.pq), len(tt.afterPush))
			}
			for i := range tt.pq {
				if tt.pq[i].Character != tt.afterPush[i].Character || tt.pq[i].Frequency != tt.afterPush[i].Frequency {
					t.Errorf("After Push(), pq[%d] = %v, expected %v", i, tt.pq[i], tt.afterPush[i])
				}
			}
		})
	}
}
