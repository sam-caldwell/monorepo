package huffman

// Len - return queue length
//
//	 Note that this is defined with the Len() name
//	 and return to satisfy the heap interface.
//
//		(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Len() int {
	return len(*pq)
}
