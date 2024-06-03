package huffman

import (
	"testing"
)

func TestPriorityQueuePush(t *testing.T) {
	var queue PriorityQueue

	node1 := Node{symbol: 'a', frequency: 3}
	node2 := Node{symbol: 'b', frequency: 2}
	node3 := Node{symbol: 'c', frequency: 10}

	t.Run("Push an element onto the priority queue", func(t *testing.T) {
		queue.Push(&node1)
		// Check if the priority queue length is correct after pushing
		if queue.Len() != 1 {
			t.Fatalf("Priority queue length after push is incorrect, got: %d, want: %d", queue.Len(), 1)
		}
	})
	t.Run("Push an element onto the priority queue", func(t *testing.T) {
		queue.Push(&node2)
		// Check if the priority queue length is correct after pushing
		if queue.Len() != 2 {
			t.Fatalf("Priority queue length after push is incorrect, got: %d, want: %d", queue.Len(), 1)
		}
	})
	t.Run("Push an element onto the priority queue", func(t *testing.T) {
		queue.Push(&node3)
		// Check if the priority queue length is correct after pushing
		if queue.Len() != 3 {
			t.Fatalf("Priority queue length after push is incorrect, got: %d, want: %d", queue.Len(), 1)
		}
	})
}
