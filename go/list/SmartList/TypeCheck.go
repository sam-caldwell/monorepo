package list

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"reflect"
)

// TypeCheck - verify that the given data is of the same type as list[0]
func (list *SmartList) TypeCheck(data any) (err error) {
	if len(*list) > 0 {
		if reflect.TypeOf((*list)[0]).Kind() != reflect.TypeOf(data).Kind() {
			err = fmt.Errorf(errors.TypeMismatch)
		}
	}
	return err
}
