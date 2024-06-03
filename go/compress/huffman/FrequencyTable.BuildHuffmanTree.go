package huffman

import "container/heap"

// BuildHuffmanTree - Given a frequency table, load the priority queue and build the huffman tree.
//
// Step 1: Load the PriorityQueue from the frequencyTable
// Step 2: Repeatedly extract the two nodes with the smallest frequencies
// Step 3: Create new parent node with these two nodes as children, and its frequency as the sum of their frequencies
// Step 4: Insert this new parent node back into the min heap
// Step 5: Continue process until there's only one node left in the heap, which becomes the root of the Huffman tree
//
//	(c) 2023 Sam Caldwell.  MIT License
func (freq *FrequencyTable) BuildHuffmanTree() *Node {
	queue := LoadPriorityQueue(freq)

	// Tree construction loop
	for queue.Len() > 1 {
		left := heap.Pop(&queue).(*Node)
		right := heap.Pop(&queue).(*Node)
		parent := &Node{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}
		heap.Push(&queue, parent)
	}
	return heap.Pop(&queue).(*Node)
}
