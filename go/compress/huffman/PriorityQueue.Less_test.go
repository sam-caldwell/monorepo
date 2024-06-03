package huffman

import (
	"testing"
)

func TestPriorityQueueLess(t *testing.T) {
	var queue PriorityQueue

	t.Run("setup data", func(t *testing.T) {
		queue = append(queue, &Node{symbol: 'a', frequency: 3})
		queue = append(queue, &Node{symbol: 'b', frequency: 2})
	})

	t.Run("check if a < b", func(t *testing.T) {
		// Check if a < b
		if !queue.Less(1, 0) {
			t.Fatalf("pq.Less(1, 0) returned false, expected true")
		}
	})

	t.Run("check if b < a (expect not)", func(t *testing.T) {
		// Check if b < a
		if queue.Less(0, 1) {
			t.Fatalf("pq.Less(0, 1) returned true, expected false")
		}
	})

}
