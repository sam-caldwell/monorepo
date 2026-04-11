package pair

// Walk - Walk the ordered pair and execute the given function on each element
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *OrderedPair[KeyType, ValueType]) Walk(fn func(key KeyType, value ValueType) error) error {
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, record := range o.data {
		if err := fn(record.Key, record.Value); err != nil {
			return err
		}
	}
	return nil
}
