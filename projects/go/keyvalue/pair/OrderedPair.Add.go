package pair

// Add - Append a
func (o *OrderedPair) Add(key string, value any) {
	o.data = append(o.data, Pair{
		Key:   key,
		Value: value,
	})
}
