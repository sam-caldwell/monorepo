package ordered

/*
 * projects/sets/ordered/Pop.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Pop - return the item at the head of the set.
func (set *Set[T]) Pop() (item T, err error) {
	if len(set.data) > 0 {
		set.lock.Lock()
		//popping item from set
		defer func() {
			defer set.lock.Unlock()
			if len(set.data) > 0 {
				set.data = set.data[1:]
			}
		}()
		item = set.data[0]
	} else {
		err = fmt.Errorf(errors.EmptySet)
	}
	return item, err
}
