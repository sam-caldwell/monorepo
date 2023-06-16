package vartype

import (
	"fmt"
	"reflect"
)

// GetFloat - Return float64 value or typecheck error
func (o *Generic) GetFloat() (data float64, err error) {
	if o.IsType(reflect.Float32) {
		return float64(o.data.(float32)), nil
	}
	if o.IsType(reflect.Float64) {
		return o.data.(float64), nil
	}
	return defaultFloatValue, fmt.Errorf(errTypeMismatch)
}
