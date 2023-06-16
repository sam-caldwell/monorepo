package simple

import (
	"fmt"
	"reflect"
)

// Add - add item to set.
func (set *Set) Add(item interface{}) error {
	if item == nil {
		return nil
	}
	if set.typ == nil {
		// Initialize the set's type with the type of the first item
		set.typ = reflect.TypeOf(item)
		set.data = make(map[interface{}]bool)
	} else if reflect.TypeOf(item) != set.typ {
		// Check if the type of the item matches the set's type
		return fmt.Errorf("item must be of type %v", reflect.TypeOf(item).String())
	}

	// Add the item to the set's data
	set.data[item] = true

	return nil
}
