package simple

import (
	"reflect"
)

// Set - Create map of any to its type
type Set struct {
	data map[any]bool
	typ  reflect.Type
}
