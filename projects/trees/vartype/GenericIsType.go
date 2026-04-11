package vartype

import "reflect"

// IsType - return boolean true if the current value is of type t
func (o *Generic) IsType(t reflect.Kind) bool {
	return reflect.TypeOf(o.data).Kind() == t
}
