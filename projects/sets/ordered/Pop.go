package ordered

/*
 * projects/sets/ordered/Pop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See README.md
 */

// Pop - return the item at the head of the set.
func (set *Set) Pop() any {
	if len(set.data) > 0 {
		set.lock.Lock()
		//popping item from set
		defer func() {
			defer set.lock.Unlock()
			if len(set.data) > 0 {
				set.data = set.data[1:]
			}
		}()
		return set.data[0]
	}
	return nil
}
