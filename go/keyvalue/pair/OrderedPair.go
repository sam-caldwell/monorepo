package pair

import "sync"

// OrderedPair - A Generic ordered key-value pair
type OrderedPair[KeyType comparable, ValueType any] struct {
	lock sync.RWMutex
	data []Pair[KeyType, ValueType]
}
