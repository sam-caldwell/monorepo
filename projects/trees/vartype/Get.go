package vartype

import (
	"fmt"
	"reflect"
)

// Get - Get a current data value, enforcing a given type of t.
func (o *Generic) Get(t reflect.Kind) (data any, err error) {
	if o.IsType(t) {
		return o.data, nil
	}
	return nil, fmt.Errorf(errTypeMismatch)
}
