package orderedset

import (
	"reflect"
)

// Set - Create map of any to its type
type Set struct {
	data []any
	typ  reflect.Type
}
