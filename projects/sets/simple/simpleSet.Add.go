package simple

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"reflect"
)

// Add - add item to set (all items must be of the same type).
func (set *Set) Add(item interface{}) error {
	if item == nil {
		return nil // We don't store nothing no how! :-)
	}
	if set.data == nil {
		// Initialize the set's type with the type of the first item
		set.data = make(map[interface{}]bool)
	}
	if len(set.data) > 0 && set.GetType() != reflect.TypeOf(item).Kind() {
		// A set must have all items of the same type as the first.
		return fmt.Errorf(exit.ErrTypeMismatch)
	}
	// Add the item to the set's data
	set.data[item] = true
	return nil
}
