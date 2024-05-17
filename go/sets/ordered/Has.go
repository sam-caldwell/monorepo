package ordered

/*
 * projects/sets/ordered/Has.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// Has scan the set and see if an item (data) has been seen before.
func (set *Set[T]) Has(data T) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return set.seenBefore(&data)
}
