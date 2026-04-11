package ordered

/*
 * projects/sets/ordered/seenBefore.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// seenBefore - returns true if an item exists
func (set *Set[T]) seenBefore(data *T) bool {
	for _, item := range set.data {
		if item == *data {
			return true
		}
	}
	return false
}
