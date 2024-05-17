package pair

func (o *OrderedPair[KeyType, ValueType]) Has(key KeyType) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	for _, element := range o.data {
		if element.Key == key {
			return true
		}
	}
	return false
}
