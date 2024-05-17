package ordered

/*
 * projects/sets/ordered
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// List - return a list of items
func (set *Set[T]) List() (result []T) {
	set.lock.RLock()
	defer set.lock.RUnlock()

	if len(set.data) > 0 {
		result = make([]T, len(set.data))
		copy(result, set.data)
	}

	return result
}
