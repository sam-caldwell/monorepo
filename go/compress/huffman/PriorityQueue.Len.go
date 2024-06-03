package huffman

// Len - return the length of the PriorityQueue
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *PriorityQueue) Len() int { return len(*pq) }
