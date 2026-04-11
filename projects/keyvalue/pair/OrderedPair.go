package pair

import "sync"

// OrderedPair - A Generic ordered key-value pair
//
//	(c) 2023 Sam Caldwell.  MIT License
type OrderedPair[KeyType comparable, ValueType any] struct {
	lock sync.RWMutex
	data []Pair[KeyType, ValueType]
}
