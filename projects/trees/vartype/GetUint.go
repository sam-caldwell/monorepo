package vartype

import (
	"fmt"
	"reflect"
)

// GetUint - return unsigned integer or typecheck
func (o *Generic) GetUint() (data uint, err error) {
	if o.IsType(reflect.Uint) {
		return o.data.(uint), nil
	}
	return defaultUintValue, fmt.Errorf(errTypeMismatch)
}
