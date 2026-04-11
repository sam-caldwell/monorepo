package ordered

/*
 * projects/sets/ordered/Count.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// Count - return a count of the records in a set
func (set *Set[T]) Count() int {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return len(set.data)
}
