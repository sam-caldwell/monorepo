package pair

// OrderedPair - A Generic ordered key-value pair
type OrderedPair[KeyType comparable, ValueType any] struct {
	data []Pair[KeyType, ValueType]
}
