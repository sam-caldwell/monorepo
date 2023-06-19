package simple

import (
	"reflect"
)

// Set - Create map of any to its type
type Set struct {
	data map[any]bool
	typ  reflect.Type
}

func (set *Set) GetFirst() any {
	if set.data == nil {
		return nil //Return nil on uninitialized set.
	}
	for k, _ := range set.data {
		return k //Exit and return the first record
	}
	return nil // Return nil on empty set
}

// GetType - Return the type of the first item (indicating type of the set)
func (set *Set) GetType() reflect.Kind {
	if v := set.GetFirst(); v != nil {
		return reflect.TypeOf(v).Kind()
	}
	return reflect.UnsafePointer
}
