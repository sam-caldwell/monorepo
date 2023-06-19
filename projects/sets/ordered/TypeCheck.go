package ordered

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"reflect"
)

// TypeCheck - verify that the given data is of the same type as the elements already there
func (set *Set) TypeCheck(data any) (err error) {
	if len(set.data) > 0 {
		set.lock.RLock()
		defer set.lock.RUnlock()
		err = set.typeCheck(&data)
	}
	return err
}

// typeCheck - the unexported / unsafe type checker
func (set *Set) typeCheck(data *any) (err error) {
	if reflect.TypeOf(set.data[0]).Kind() == reflect.TypeOf(data).Kind() {
		return nil
	}
	return fmt.Errorf(exit.ErrTypeMismatch)
}
