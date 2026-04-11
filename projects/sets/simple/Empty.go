package simple

// Empty - Return whether the set is empty.
func (set *Set[T]) Empty() bool {
	return set.Count() == 0
}
