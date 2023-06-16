package vartype

import (
	"fmt"
	"reflect"
)

// GetInt32 - return integer or typecheck error
func (o *Generic) GetInt32() (data int32, err error) {
	if o.IsType(reflect.Int32) {
		return o.data.(int32), nil
	}
	return defaultIntValue, fmt.Errorf(errTypeMismatch)
}
