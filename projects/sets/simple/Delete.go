package simple

// Delete - delete an item from the set if it exists (or noop if it doesn't exist).
func (set *Set[T]) Delete(item T) {
	delete(set.data, item)
}
