package vartype

import (
	"fmt"
	"reflect"
)

// GetInt - return integer or typecheck error
func (o *Generic) GetInt() (data int, err error) {
	if o.IsType(reflect.Int) {
		return o.data.(int), nil
	}
	return defaultIntValue, fmt.Errorf(errTypeMismatch)
}
