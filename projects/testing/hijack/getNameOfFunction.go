package hijack

import (
	"reflect"
	"runtime"
)

// getNameOfFunction - Return the name of the symbol 'arbitraryObject'
func getNameOfFunction(arbitraryObject any) string {
	return runtime.FuncForPC(
		reflect.ValueOf(
			arbitraryObject).
			Pointer()).
		Name()
}
