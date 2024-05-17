package pair

// Walk - Walk the ordered pair and execute the given function on each element
func (o *OrderedPair[KeyType, ValueType]) Walk(fn func(key KeyType, value ValueType) error) error {
	for _, record := range o.data {
		if err := fn(record.Key, record.Value); err != nil {
			return err
		}
	}
	return nil
}
