package simple

import "reflect"

// GetType - Return the type of the first item (indicating type of the set)
func (set *Set) GetType() reflect.Kind {
	if v := set.GetFirst(); v != nil {
		return reflect.TypeOf(v).Kind()
	}
	return reflect.UnsafePointer
}
