package pair

func (o *OrderedPair) Has(key string) bool {
	for _, element := range o.data {
		if element.Key == key {
			return true
		}
	}
	return false
}
