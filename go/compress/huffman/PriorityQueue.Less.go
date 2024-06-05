package huffman

// Less - compare two elements of the PriorityQueue and return boolean if a < b
//
//	 We want the lower frequency to be the higher priority
//	 But where our frequency is the same we will break the tie by comparing the symbol.
//
//		(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Less(a, b int) bool {
	if (*pq)[a].frequency == (*pq)[b].frequency {
		return (*pq)[a].symbol < (*pq)[b].symbol
	}
	return (*pq)[a].frequency < (*pq)[b].frequency
}
