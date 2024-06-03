package huffman

import (
	"container/heap"
	"testing"
)

func TestPriorityQueuePush(t *testing.T) {
	// Create a sample priority queue
	pq := make(PriorityQueue, 0)

	// Push an element onto the priority queue
	heap.Push(&pq, &Node{symbol: 'a', frequency: 3})

	// Check if the priority queue length is correct after pushing
	if pq.Len() != 1 {
		t.Fatalf("Priority queue length after push is incorrect, got: %d, want: %d", pq.Len(), 1)
	}

	// Peek the top element from the priority queue
	top := pq[0]

	// Check if the top element has the correct values
	if top.symbol != 'a' || top.frequency != 3 {
		t.Fatalf("Top element after push has incorrect values, got: %v, want: %v", top, &Node{symbol: 'a', frequency: 3})
	}
}
