package pair

// Add - Append a key-value pair to the ordered pair
func (o *OrderedPair[KeyType, ValueType]) Add(key KeyType, value ValueType) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.data = append(o.data, Pair[KeyType, ValueType]{
		Key:   key,
		Value: value,
	})
}
