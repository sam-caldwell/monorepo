package huffman

// Less - return bool if i less than j
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].Frequency < (*pq)[j].Frequency
}
