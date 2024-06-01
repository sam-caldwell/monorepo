package huffman

import "testing"

func TestPriorityQueue_map(t *testing.T) {
	t.Run("default state", func(t *testing.T) {
		var p priorityQueue
		if len(p) != 0 {
			t.Fatal("priority queue has non-zero length")
		}
		p = append(p, &node{})
		if len(p) != 1 {
			t.Fatal("priority queue expects length 1")
		}
	})
}
