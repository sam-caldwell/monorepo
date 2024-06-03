package huffman

import (
	"reflect"
	"testing"
)

// Test traverseTree function
func TestTraverseTree(t *testing.T) {
	// Edge Case 1: Empty Tree
	t.Run("Empty Tree", func(t *testing.T) {
		var root *Node
		expectedCodes := make(CodeMap)
		codes := make(CodeMap)
		traverseTree(root, []byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			t.Fatalf("traverseTree() for empty tree returned %v, expected %v", codes, expectedCodes)
		}
	})

	// Edge Case 2: Tree with only one node
	t.Run("Tree with one node", func(t *testing.T) {
		root := &Node{
			frequency: 10,
			symbol:    'a',
		}
		expectedCodes := CodeMap{'a': []byte("0")}
		codes := make(CodeMap)
		traverseTree(root, []byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for tree with one node returned %v, expected %v", codes, expectedCodes)
		}
	})

	// Negative Test 1: Nil Tree
	t.Run("nil Tree", func(t *testing.T) {
		var root *Node
		codes := make(CodeMap)
		traverseTree(root, []byte(""), codes)
		if len(codes) != 0 {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for nil tree returned %v, expected empty map", codes)
		}
	})

	// Negative Test 2: Invalid Input
	t.Run("Invalid Input", func(t *testing.T) {
		root := &Node{
			frequency: 10,
			symbol:    'a',
			left:      &Node{},
			right:     nil,
		}
		codes := make(CodeMap)
		traverseTree(root, []byte(""), codes)
		if len(codes) != 0 {
			root.PrettyPrint()
			t.Fatalf("traverseTree() for tree with invalid input returned %v, expected empty map", codes)
		}
	})

	// Test with a sample Huffman tree
	t.Run("Sample Huffman Tree", func(t *testing.T) {
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
			'a': []byte("010"),
			'b': []byte("011"),
			'c': []byte("000"),
			'd': []byte("001"),
			'e': []byte("10"),
			'f': []byte("1"),
		}

		codes := make(CodeMap)
		traverseTree(root, []byte(""), codes)
		if !reflect.DeepEqual(codes, expectedCodes) {
			root.PrettyPrint()
			t.Fatalf("traverseTree() returned %v, expected %v", codes, expectedCodes)
		}
	})
}
