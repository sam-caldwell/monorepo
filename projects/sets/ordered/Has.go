package ordered

// Has scan the set and see if an item (data) has been seen before.
func (set *Set) Has(data any) bool {
	if data == nil {
		return false
	}

	set.lock.RLock()
	defer set.lock.RUnlock()

	return set.seenBefore(&data)
}

// seenBefore - returns true if an item exists
func (set *Set) seenBefore(data *any) bool {
	for _, item := range set.data {
		if item == *data {
			return true
		}
	}
	return false
}
