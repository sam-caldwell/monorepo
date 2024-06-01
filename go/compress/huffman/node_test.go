package huffman

import "testing"

func TestNode_struct(t *testing.T) {
	t.Run("default state", func(t *testing.T) {
		var n node
		if n.Character != 0x00 {
			t.Fatal("expected nil character")
		}
		if n.Frequency != 0 {
			t.Fatal("expected 0 frequency")
		}
		if n.Left != nil {
			t.Fatal("expected nil left")
		}
		if n.Right != nil {
			t.Fatal("expected nil right")
		}
	})
}
