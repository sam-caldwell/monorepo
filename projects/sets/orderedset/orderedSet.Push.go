package orderedset

// Push - Push the item to the top of the list (n+1)
func (set *Set) Push(item any) error {
	return set.Add(item)
}
