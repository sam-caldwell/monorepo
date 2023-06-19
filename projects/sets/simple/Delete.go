package simple

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
)

// Delete - delete an item from the set
func (set *Set) Delete(item interface{}) error {
	if set.data == nil {
		return fmt.Errorf(exit.ErrNotInitialized)
	}

	if !set.Has(item) {
		return fmt.Errorf(exit.ErrNotFound)
	}

	delete(set.data, item)
	return nil
}
