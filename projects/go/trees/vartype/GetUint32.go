package vartype

import (
	"fmt"
	"reflect"
)

// GetUint32 - return unsigned integer or typecheck
func (o *Generic) GetUint32() (data uint32, err error) {
	if o.IsType(reflect.Uint32) {
		return o.data.(uint32), nil
	}
	return defaultUintValue, fmt.Errorf(errTypeMismatch)
}
