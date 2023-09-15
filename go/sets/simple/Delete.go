package simple

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Delete - delete an item from the set
func (set *Set) Delete(item interface{}) error {
	if set.data == nil {
		return fmt.Errorf(errors.NotInitialized)
	}

	if !set.Has(item) {
		return fmt.Errorf(errors.NotFound)
	}

	delete(set.data, item)
	return nil
}
