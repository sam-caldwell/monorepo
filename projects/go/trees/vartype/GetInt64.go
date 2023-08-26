package vartype

import (
	"fmt"
	"reflect"
)

// GetInt64 - return integer or typecheck error
func (o *Generic) GetInt64() (data int64, err error) {
	if o.IsType(reflect.Int64) {
		return o.data.(int64), nil
	}
	return defaultIntValue, fmt.Errorf(errTypeMismatch)
}
