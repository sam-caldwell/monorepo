package huffman

import (
	"fmt"
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	t.Run("Edge Case 1: Empty Input", func(t *testing.T) {
		frequencyTable := FrequencyTable{}
		expectedFrequency := uint(0)
		var root *Node
		root = frequencyTable.BuildHuffmanTree()
		if root.frequency != expectedFrequency {
			t.Errorf("Root node has incorrect frequency, got: %d, want: %d", root.frequency, expectedFrequency)
		}
		if err := ValidateTreeStructure(root); err != nil {
			root.PrettyPrint()
			t.Fatal(err)
		}
	})
	t.Run("Edge Case 2: Input with one symbol", func(t *testing.T) {
		frequencyTable := FrequencyTable{
			'a': 1,
		}
		expectedFrequency := uint(1)
		var root *Node
		root = frequencyTable.BuildHuffmanTree()
		if root.frequency != expectedFrequency {
			t.Errorf("Root node has incorrect frequency, got: %d, want: %d", root.frequency, expectedFrequency)
		}
		if err := ValidateTreeStructure(root); err != nil {
			root.PrettyPrint()
			t.Fatal(err)
		}
	})
	t.Run("Normal Case 1: Input with multiple symbols", func(t *testing.T) {
		frequencyTable := FrequencyTable{
			'a': 5,
			'b': 9,
			'c': 12,
			'd': 13,
			'e': 16,
			'f': 45,
		}
		var root *Node

		t.Run("Build the Huffman tree", func(t *testing.T) {
			root = frequencyTable.BuildHuffmanTree()
		})
		t.Run("Verify the node's frequency equal to the sum of all frequencies", func(t *testing.T) {
			expectedFrequency := uint(100) // The sum of all frequencies in the frequency table
			if root.frequency != expectedFrequency {
				t.Errorf("Root node has incorrect frequency, got: %d, want: %d", root.frequency, expectedFrequency)
			}
		})
		t.Run("Verify the root node has two children", func(t *testing.T) {
			// Check if the root node has two children
			if root.left == nil || root.right == nil {
				t.Errorf("Root node does not have two children")
			}
		})
		t.Run("Analyze the tree structure", func(t *testing.T) {
			root.PrettyPrint()
			if err := ValidateTreeStructure(root); err != nil {
				t.Fatal(err)
			}
		})
	})
}

// ValidateTreeStructure traverses the Huffman tree and validates the structure
func ValidateTreeStructure(node *Node) error {
	if node.left == nil && node.right == nil {
		return nil // Leaf node
	}
	if node.left != nil {
		// Check left subtree
		if node.left.frequency > node.frequency {
			return fmt.Errorf("left child has higher frequency than parent, got: %d, parent: %d",
				node.left.frequency, node.frequency)
		}
		if err := ValidateTreeStructure(node.left); err != nil {
			return err
		}
	}
	if node.right != nil {
		// Check right subtree
		if node.right.frequency > node.frequency {
			return fmt.Errorf("right child has higher frequency than parent, got: %d, parent: %d",
				node.right.frequency, node.frequency)
		}
		if err := ValidateTreeStructure(node.right); err != nil {
			return err
		}
	}
	return nil
}
