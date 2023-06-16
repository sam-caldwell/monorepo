package ordered

import (
	"reflect"
)

// Set - Create map of any to its type
type Set struct {
	data []any
	typ  reflect.Type
}
