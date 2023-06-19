package ordered

// Empty - Return boolean true if set has 1 or more records
func (set *Set) Empty() bool {
	return set.Count() > 0
}
