package vartype

import "reflect"

// Type - return the actual underlying type.
func (o *Generic) Type() (t reflect.Kind) {
	return reflect.TypeOf(o.data).Kind()
}
