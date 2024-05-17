package pair

// Pair - A generic key-value pair element (used independently or with OrderedPair
type Pair[KeyType comparable, ValueType any] struct {
	Key   KeyType
	Value ValueType
}
