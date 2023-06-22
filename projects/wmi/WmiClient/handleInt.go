package wmiclient

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

// handleInt - Cast int64 into the given field
func handleInt(field reflect.Value, value int64) (err error) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		field.SetInt(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		field.SetUint(uint64(value))
	default:
		err = fmt.Errorf(errors.TypeMismatchWithDetail, field.Kind().String())
	}
	return err
}
