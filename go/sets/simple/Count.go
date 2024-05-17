package simple

// Count - return a count of the records in a set
func (set *Set[T]) Count() int {
	if set.data == nil {
		return 0
	}
	return len(set.data)
}
