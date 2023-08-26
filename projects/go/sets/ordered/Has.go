package ordered

/*
 * projects/sets/ordered/Has.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See README.md
 */

// Has scan the set and see if an item (data) has been seen before.
func (set *Set) Has(data any) bool {
	if data == nil {
		return false
	}

	set.lock.RLock()
	defer set.lock.RUnlock()

	return set.seenBefore(&data)
}
