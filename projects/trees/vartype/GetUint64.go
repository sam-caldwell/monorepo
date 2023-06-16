package vartype

import (
	"fmt"
	"reflect"
)

// GetUint64 - return unsigned integer or typecheck
func (o *Generic) GetUint64() (data uint64, err error) {
	if o.IsType(reflect.Uint64) {
		return o.data.(uint64), nil
	}
	return defaultUintValue, fmt.Errorf(errTypeMismatch)
}
