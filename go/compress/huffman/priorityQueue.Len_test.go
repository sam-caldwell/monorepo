package huffman

import "testing"

func TestPriorityQueue_Len(t *testing.T) {
	var p priorityQueue
	t.Run("default state", func(t *testing.T) {
		if p.Len() != 0 {
			t.Errorf("p.Len() = %d, want 0", p.Len())
		}
	})
	t.Run("add item to queue1", func(t *testing.T) {
		p = append(p, &(node{}))
	})
	t.Run("verify size 1", func(t *testing.T) {
		if p.Len() != 1 {
			t.Errorf("p.Len() = %d, want 1", p.Len())
		}
	})
	t.Run("add item to queue2", func(t *testing.T) {
		p = append(p, &(node{}))
	})
	t.Run("verify size 2", func(t *testing.T) {
		if p.Len() != 2 {
			t.Errorf("p.Len() = %d, want 2", p.Len())
		}
	})
}
