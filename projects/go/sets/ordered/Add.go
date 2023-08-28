package ordered

/*
 * projects/sets/ordered/Add.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file adds an arbitrary-typed object to the
 * ordered set and enforces type-checking if data
 * exists already.  First element determines type.
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit/errors"
)

// Add - add item to set if the item is the same type as the set
func (set *Set) Add(item any) (err error) {
	//Bail on nil inputs.  not worth the time.
	if item == nil {
		return err
	}

	set.lock.Lock()
	defer set.lock.Unlock()

	// If there is a first element, perform our typecheck.
	if len(set.data) > 0 {
		if err = set.typeCheck(&item); err != nil {
			return err
		}
	}

	// Make sure we don't store duplicates.
	if set.seenBefore(&item) {
		return fmt.Errorf(errors.DuplicateEntry)
	}

	// Add the item to the set.
	set.data = append(set.data, item)
	return err
}

// insert - a simple (unexported insert method)
func (set *Set) insert(item any) {
	set.data = append(set.data, item)
}
