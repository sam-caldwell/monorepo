package huffman

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	var queue PriorityQueue
	if len(queue) != 0 {
		t.Fatal("Priority Queue should be empty")
	}
	queue = append(queue, nil)
	if len(queue) != 1 {

	}
	if queue[0] != nil {
		t.Fatal("Priority Queue element 0 should be nil")
	}
	n := Node{}
	queue[0] = &(n)
	if !reflect.DeepEqual(n, *queue[0]) {
		t.Fatal("Priority Queue element not expected")
	}
}
