package vartype

import (
	"fmt"
	"reflect"
)

// GetFloat64 - return float64 or typecheck error
func (o *Generic) GetFloat64() (data float64, err error) {
	if o.IsType(reflect.Float64) {
		return o.data.(float64), nil
	}
	return defaultFloatValue, fmt.Errorf(errTypeMismatch)
}
