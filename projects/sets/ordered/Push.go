package ordered

/*
 * projects/sets/ordered/Push.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

// Push - Push the item to the top of the list (n+1)
func (set *Set[T]) Push(item T) error {
	return set.Add(item)
}
