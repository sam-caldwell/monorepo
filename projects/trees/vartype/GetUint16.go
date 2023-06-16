package vartype

import (
	"fmt"
	"reflect"
)

// GetUint16 - return unsigned integer or typecheck
func (o *Generic) GetUint16() (data uint16, err error) {
	if o.IsType(reflect.Uint16) {
		return o.data.(uint16), nil
	}
	return defaultUintValue, fmt.Errorf(errTypeMismatch)
}
