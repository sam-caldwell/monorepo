package simple

import "reflect"

// TypeCheck - Return whether item is of the same type as the set (empty defaults to true)
func (set *Set) TypeCheck(item any) bool {
	if set.Empty() {
		return true
	}
	return set.GetType() == reflect.TypeOf(item).Kind()
}
