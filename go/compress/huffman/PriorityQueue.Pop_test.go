package huffman

import (
	"reflect"
	"testing"
)

func TestPriorityQueuePop(t *testing.T) {
	var queue PriorityQueue

	node1 := Node{symbol: 'a', frequency: 3}
	node2 := Node{symbol: 'b', frequency: 2}
	node3 := Node{symbol: 'c', frequency: 10}

	t.Run("Add some elements to the priority queue", func(t *testing.T) {
		queue = append(queue, &node1)
		queue = append(queue, &node2)
		queue = append(queue, &node3)
	})
	t.Run("Pop element3 from the priority queue", func(t *testing.T) {
		item := queue.Pop().(*Node)
		if !reflect.DeepEqual(*item, node3) {
			t.Fatalf("Popped element has incorrect values, got: %v, want: %v", *item, node3)
		}
	})
	t.Run("Pop element2 from the priority queue", func(t *testing.T) {
		item := queue.Pop().(*Node)
		if !reflect.DeepEqual(*item, node2) {
			t.Fatalf("Popped element has incorrect values, got: %v, want: %v", *item, node2)
		}
	})
	t.Run("Pop element1 from the priority queue", func(t *testing.T) {
		item := queue.Pop().(*Node)
		if !reflect.DeepEqual(*item, node1) {
			t.Fatalf("Popped element has incorrect values, got: %v, want: %v", *item, node1)
		}
	})
	t.Run("Ensure the priority queue is empty", func(t *testing.T) {
		if sz := queue.Len(); sz != 0 {
			t.Fatalf("Priority queue is not empty after popping all elements, got length %d", sz)
		}
	})
}
