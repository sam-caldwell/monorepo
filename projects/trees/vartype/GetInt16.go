package vartype

import (
	"fmt"
	"reflect"
)

// GetInt16 - return integer or typecheck error
func (o *Generic) GetInt16() (data int16, err error) {
	if o.IsType(reflect.Int16) {
		return o.data.(int16), nil
	}
	return defaultIntValue, fmt.Errorf(errTypeMismatch)
}
