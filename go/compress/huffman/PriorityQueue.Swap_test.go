package huffman

import (
	"reflect"
	"testing"
)

func TestPriorityQueueSwap(t *testing.T) {
	var queue PriorityQueue

	node0 := Node{symbol: 'a', frequency: 3}
	node1 := Node{symbol: 'b', frequency: 2}
	node2 := Node{symbol: 'c', frequency: 10}

	t.Run("Add elements to the queue", func(t *testing.T) {
		queue = append(queue, &node0)
		queue = append(queue, &node1)
		queue = append(queue, &node2)
	})
	t.Run("Verify element position", func(t *testing.T) {
		if p := queue[0]; !reflect.DeepEqual(node0, *p) {
			t.Fatalf("element0 not equal.  Got %v.  Expected %v", *p, node0)
		}
		if p := queue[1]; !reflect.DeepEqual(node1, *p) {
			t.Fatalf("element1 not equal.  Got %v.  Expected %v", *p, node1)
		}
		if p := queue[2]; !reflect.DeepEqual(node2, *p) {
			t.Fatalf("element2 not equal.  Got %v.  Expected %v", *p, node2)
		}
	})
	t.Run("Swap the first and last elements", func(t *testing.T) {
		queue.Swap(0, 2)
	})
	t.Run("Verify element position", func(t *testing.T) {
		if p := queue[0]; !reflect.DeepEqual(node2, *p) {
			t.Fatalf("element0 not equal.  Got %v.  Expected %v", *p, node2)
		}
		if p := queue[1]; !reflect.DeepEqual(node1, *p) {
			t.Fatalf("element1 not equal.  Got %v.  Expected %v", *p, node1)
		}
		if p := queue[2]; !reflect.DeepEqual(node0, *p) {
			t.Fatalf("element2 not equal.  Got %v.  Expected %v", *p, node0)
		}
	})
}
