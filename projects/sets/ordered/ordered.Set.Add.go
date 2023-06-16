package ordered

import (
	"fmt"
	"reflect"
)

// Add - add item to set.
func (set *Set) Add(item interface{}) error {
	if item == nil {
		return nil
	}
	if len(set.data) == 0 {
		set.typ = reflect.TypeOf(item)
	} else if reflect.TypeOf(set.data[0]) != reflect.TypeOf(item) {
		return fmt.Errorf("item must be of type %v", reflect.TypeOf(item).String())
	}
	set.data = append(set.data, item)
	return nil
}
