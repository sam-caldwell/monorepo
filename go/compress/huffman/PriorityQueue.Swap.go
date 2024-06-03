package huffman

// Swap - Given two nodes in PriorityQueue, swap elements a and b.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Swap(a, b int) {
	//(*pq)[a], (*pq)[b] = (*pq)[b], (*pq)[a]
	//
	tmp := (*pq)[a]
	(*pq)[a] = (*pq)[b]
	(*pq)[b] = tmp
}
