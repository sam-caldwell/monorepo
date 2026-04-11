package vartype

import (
	"fmt"
	"reflect"
)

// GetBool - Return a boolean value or type mismatch error
func (o *Generic) GetBool() (data bool, err error) {
	if o.IsType(reflect.Bool) {
		return o.data.(bool), nil
	}
	return defaultBoolValue, fmt.Errorf(errTypeMismatch)

}
