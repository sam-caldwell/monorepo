package vartype

import (
	"fmt"
	"reflect"
)

// GetString - return the string data or typecheck error
func (o *Generic) GetString() (data string, err error) {
	if o.IsType(reflect.String) {
		return o.data.(string), nil
	}
	return defaultStringValue, fmt.Errorf(errTypeMismatch)
}
