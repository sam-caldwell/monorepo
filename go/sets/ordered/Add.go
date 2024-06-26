package ordered

/*
 * projects/sets/ordered/Add.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds an arbitrary-typed object to the
 * ordered set.  Returns error on duplicate entry.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Add - add item to set if the item is the same type as the set
func (set *Set[T]) Add(item T) (err error) {
	set.lock.Lock()
	defer set.lock.Unlock()

	// Make sure we don't store duplicates.
	if set.seenBefore(&item) {
		return fmt.Errorf(errors.DuplicateEntry)
	}

	// Add the item to the set.
	set.data = append(set.data, item)
	return err
}
