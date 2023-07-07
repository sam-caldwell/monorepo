package hijack

import (
	"fmt"
	"reflect"
)

// requireCompatibleFunctions - Return an error if the given function is not a function.
func requireCompatibleFunctions(a, b any) (err error) {

	const (
		compatibilityError = "both %s and %s must be of the same type"
		requiredFuncError  = "%s has to be a function"
	)

	if reflect.ValueOf(a).Kind() != reflect.Func {
		err = fmt.Errorf(requiredFuncError, getNameOfFunction(a))
	}
	if reflect.ValueOf(b).Kind() != reflect.Func {
		err = fmt.Errorf(requiredFuncError, getNameOfFunction(b))
	}
	if reflect.ValueOf(a).Type() != reflect.ValueOf(b).Type() {
		err = fmt.Errorf(compatibilityError, getNameOfFunction(a), getNameOfFunction(b))
	}
	return err
}
