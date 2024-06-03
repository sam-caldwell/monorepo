package huffman

import "testing"

func TestNode_struct(t *testing.T) {
	var n = Node{
		symbol:    0xFF,
		frequency: 65534,
		left:      &Node{},
		right:     &Node{},
	}
	if n.symbol != 0xFF {
		t.Fatalf("symbol mismatch")
	}
	if n.frequency != 65534 {
		t.Fatalf("frequency mismatch")
	}
	if n.left.symbol != 0x00 {
		t.Fatalf("left symbol mismatch")
	}
	if n.right.symbol != 0x00 {
		t.Fatalf("right symbol mismatch")
	}
	n.left = nil
	n.right = nil
	if n.left != nil || n.right != nil {
		t.Fatalf("left and right should be nil")
	}
}
