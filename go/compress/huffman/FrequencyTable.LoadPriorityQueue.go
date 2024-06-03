package huffman

import "container/heap"

// LoadPriorityQueue - Create, initialize and load the priority queue from the frequency table.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (freq *FrequencyTable) LoadPriorityQueue() PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for symbol, frequency := range *freq {
		heap.Push(&pq, &Node{symbol: symbol, frequency: frequency})
	}
	return pq
}
