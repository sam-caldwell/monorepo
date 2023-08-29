package vartype

import (
	"fmt"
	"reflect"
)

// GetUint8 - return unsigned integer or typecheck
func (o *Generic) GetUint8() (data uint8, err error) {
	if o.IsType(reflect.Uint8) {
		return o.data.(uint8), nil
	}
	return defaultUintValue, fmt.Errorf(errTypeMismatch)
}
