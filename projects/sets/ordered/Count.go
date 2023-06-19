package ordered

// Count - return a count of the records in a set
func (set *Set) Count() int {
	return len(set.data)
}
