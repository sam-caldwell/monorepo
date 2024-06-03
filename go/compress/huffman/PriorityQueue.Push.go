package huffman

// Push - push an element (node) to the PriorityQueue
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Push(element any) {
	// Push the element onto the heap
	*pq = append(*pq, element.(*Node))
}
