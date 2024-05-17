package simple

func (set *Set[T]) Empty() bool {
	return set.Count() == 0
}
