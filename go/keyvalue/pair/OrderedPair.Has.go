package pair

// Has - Return boolean value indicating if ordered pair set has the given key.
//
//	     Note that this is a simple linear search and performs accordingly.
//	     for larger sets, use a different search mechanism.
//
//		(c) 2023 Sam Caldwell.  MIT License
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
