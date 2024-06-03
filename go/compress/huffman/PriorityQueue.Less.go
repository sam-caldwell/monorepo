package huffman

// Less - compare two elements of the PriorityQueue and return boolean if a < b
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Less(a, b int) bool {
	// We want the lower frequency to be the higher priority
	return (*pq)[a].frequency < (*pq)[b].frequency
}
