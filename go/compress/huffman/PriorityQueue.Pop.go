package huffman

// Pop - Pop an element from the PriorityQueue
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
