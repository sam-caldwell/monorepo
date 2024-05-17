package simple

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// GetFirst - Return the first element (not guaranteed ordering)
func (set *Set[T]) GetFirst() (item T, err error) {

	if set.data != nil {

		for k := range set.data {
			return k, nil //Exit and return the first record
		}

	}
	return item, fmt.Errorf(errors.EmptySet)
}
