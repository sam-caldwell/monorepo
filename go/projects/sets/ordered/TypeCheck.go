package ordered

/*
 * projects/sets/ordered/Set.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * See README.md
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/exit/errors"
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
	if reflect.TypeOf(set.data[0]).Kind().String() != reflect.TypeOf(*data).Kind().String() {
		return fmt.Errorf(errors.TypeMismatch)
	}
	return nil
}
