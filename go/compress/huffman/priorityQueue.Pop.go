package huffman

// Pop - pop and return priority queue element
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	if n == 0 {
		return &(node{})
	}
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
