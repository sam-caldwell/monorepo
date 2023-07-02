package ordered

/*
 * projects/sets/ordered/Add.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds an arbitrary-typed object to the
 * ordered set and enforces type-checking.
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

// Add - add item to set if the item is the same type as the set
func (set *Set) Add(item any) (err error) {
	//Bail on nil inputs.  not worth the time.
	if item == nil {
		return err
	}

	set.lock.Lock()
	defer set.lock.Unlock()

	if len(set.data) > 0 {
		if err = set.typeCheck(&item); err != nil {
			return err
		}
	}
	if set.seenBefore(&item) {
		return fmt.Errorf(errors.DuplicateEntry)
	}

	set.data = append(set.data, item)
	return err
}

// insert - a simple (unexported insert method)
func (set *Set) insert(item any) {
	set.data = append(set.data, item)
}
