package huffman

import (
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Pop - pop and return priority queue element
//
//	(c) 2023 Sam Caldwell.  MIT License
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	if n == 0 {
		panic(errors.EmptySet)
	}
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}
