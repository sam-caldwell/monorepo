package huffman

// Swap - swap element i and j
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}
