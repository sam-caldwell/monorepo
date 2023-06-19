package simple

func (set *Set) GetFirst() any {
	if set.data == nil {
		return nil //Return nil on uninitialized set.
	}
	for k := range set.data {
		return k //Exit and return the first record
	}
	return nil // Return nil on empty set
}
