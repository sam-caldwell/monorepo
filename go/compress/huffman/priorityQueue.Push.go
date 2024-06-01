package huffman

// Push - push element to queue
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Push(x any) {
	node := x.(*node)
	*pq = append(*pq, node)
}
