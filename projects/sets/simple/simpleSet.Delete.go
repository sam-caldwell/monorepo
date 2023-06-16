package simpleset

import "fmt"

// Delete - delete an item from the set
func (set *Set) Delete(item interface{}) error {
	if set.data == nil {
		return fmt.Errorf("set is not initialized")
	}

	if !set.Has(item) {
		return fmt.Errorf("item does not exist in the set")
	}

	delete(set.data, item)
	return nil
}
