package pair

// Pair - A generic key-value pair element (used independently or with OrderedPair
//
//	(c) 2023 Sam Caldwell.  MIT License
type Pair[KeyType comparable, ValueType any] struct {
	Key   KeyType
	Value ValueType
}
