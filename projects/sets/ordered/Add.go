package ordered

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

func (set *Set) insert(item any) {
	set.data = append(set.data, item)
}
