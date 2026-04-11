package vartype

import (
	"fmt"
	"reflect"
)

// GetFloat32 - return a float32 object or typecheck error
func (o *Generic) GetFloat32() (data float32, err error) {
	if o.IsType(reflect.Float32) {
		return o.data.(float32), nil
	}
	return defaultFloatValue, fmt.Errorf(errTypeMismatch)
}
