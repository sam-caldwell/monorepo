package huffman

import "container/heap"

// buildHuffmanTree - build the huffman tree
//
//	(c) 2023 Sam Caldwell.  MIT License
func buildHuffmanTree(input []byte) *node {
	// Calculate frequency of each character
	frequency := make(map[byte]int)
	for _, b := range input {
		frequency[b]++
	}

	// Create a priority queue and add all characters with their frequencies
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	for char, freq := range frequency {
		heap.Push(&pq, &node{Character: char, Frequency: freq})
	}

	// Build the Huffman tree
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*node)
		right := heap.Pop(&pq).(*node)
		merged := &node{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}
		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*node)
}
