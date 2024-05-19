package pair

// Count - Return the size of the OrderedPair
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *OrderedPair[KeyType, ValueType]) Count() int {
	return len(o.data)
}
