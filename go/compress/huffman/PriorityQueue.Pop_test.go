package huffman

import (
	"container/heap"
	"testing"
)

func TestPriorityQueuePop(t *testing.T) {
	// Create a sample priority queue
	pq := make(PriorityQueue, 0)

	// Add some elements to the priority queue
	heap.Push(&pq, &Node{symbol: 'a', frequency: 3})
	heap.Push(&pq, &Node{symbol: 'b', frequency: 2})

	// Pop an element from the priority queue
	item := pq.Pop().(*Node)

	// Check if the correct element was popped
	if item.symbol != 'a' || item.frequency != 3 {
		t.Fatalf("Popped element has incorrect values, got: %v, want: %v", item, &Node{symbol: 'a', frequency: 3})
	}

	// Pop another element from the priority queue
	item = pq.Pop().(*Node)

	// Check if the correct element was popped
	if item.symbol != 'b' || item.frequency != 2 {
		t.Fatalf("Popped element has incorrect values, got: %v, want: %v", item, &Node{symbol: 'b', frequency: 2})
	}

	// Ensure the priority queue is empty
	if pq.Len() != 0 {
		t.Fatalf("Priority queue is not empty after popping all elements, got length %d", pq.Len())
	}
}
