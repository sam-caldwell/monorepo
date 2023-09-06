package simple

func (set *Set) Empty() bool {
	return set.Count() <= 0
}
