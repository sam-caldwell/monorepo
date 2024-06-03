package huffman

import "testing"

func TestPriorityQueue_Len(t *testing.T) {
	var queue PriorityQueue
	if queue.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", queue.Len())
	}
	queue = append(queue, nil)
	if queue.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", queue.Len())
	}
	queue = append(queue, nil)
	if queue.Len() != 2 {
		t.Fatalf("Len() = %d, want 2", queue.Len())
	}
	queue = append(queue, nil)
	if queue.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", queue.Len())
	}
	queue = append(queue, nil)
	if queue.Len() != 4 {
		t.Fatalf("Len() = %d, want 4", queue.Len())
	}
}
