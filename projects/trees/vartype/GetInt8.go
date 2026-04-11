package vartype

import (
	"fmt"
	"reflect"
)

// GetInt8 - return integer or typecheck error
func (o *Generic) GetInt8() (data int8, err error) {
	if o.IsType(reflect.Int8) {
		return o.data.(int8), nil
	}
	return defaultIntValue, fmt.Errorf(errTypeMismatch)
}
