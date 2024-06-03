package huffman

import (
	"reflect"
	"testing"
)

// Test traverseTree function
func TestTraverseTree(t *testing.T) {
	var root *Node
	t.Run("Edge Case 1: Empty Tree", func(t *testing.T) {
		expectedCodes := make(CodeMap)
		codes := make(CodeMap)
		root.traverseTree([]byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Fatalf("traverseTree() for empty tree returned %v, expected %v", codes, expectedCodes)
		}
	})
	t.Run("Edge Case 2: Tree with only one node", func(t *testing.T) {
		root = &Node{
			frequency: 10,
			symbol:    'a',
		}
		expectedCodes := CodeMap{'a': []byte{}}
		codes := make(CodeMap)
		root.traverseTree([]byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for tree with one node returned %v, expected %v", codes, expectedCodes)
		}
	})
	t.Run("Negative Test 1: Nil Tree", func(t *testing.T) {
		var root *Node
		codes := make(CodeMap)
		root.traverseTree([]byte(""), codes)
		if len(codes) != 0 {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for nil tree returned %v, expected empty map", codes)
		}
	})
	t.Run("Negative Test 2: Invalid Input", func(t *testing.T) {
		root := &Node{
			frequency: 10,
			symbol:    'a',
			left:      &Node{},
			right:     nil,
		}
		codes := make(CodeMap)
		root.traverseTree([]byte(""), codes)
		if len(codes) != 1 {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for tree with invalid input returned %v, expected empty map", codes)
		}
	})
	t.Run("Test with a sample Huffman tree", func(t *testing.T) {
		root := &Node{
			frequency: 100,
			left: &Node{
				frequency: 45,
				symbol:    'f',
			},
			right: &Node{
				frequency: 55,
				left: &Node{
					frequency: 25,
					left: &Node{
						frequency: 12,
						symbol:    'c',
					},
					right: &Node{
						frequency: 13,
						symbol:    'd',
					},
				},
				right: &Node{
					frequency: 30,
					left: &Node{
						frequency: 14,
						left: &Node{
							frequency: 5,
							symbol:    'a',
						},
						right: &Node{
							frequency: 9,
							symbol:    'b',
						},
					},
					right: &Node{
						frequency: 16,
						symbol:    'e',
					},
				},
			},
		}
		expectedCodes := CodeMap{
			'a': []byte{1, 1, 0, 0},
			'b': []byte{1, 1, 0, 1},
			'c': []byte{1, 0, 0},
			'd': []byte{1, 0, 1},
			'e': []byte{1, 1, 1},
			'f': []byte{0},
		}

		codes := make(CodeMap)
		root.traverseTree([]byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			root.PrettyPrint()
			t.Fatalf("traverseTree()\nreturned %v,\nexpected %v",
				codes, expectedCodes)
		}
	})
}
