package wmiclient

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

// handleString - Given a string value, cast it into a compatible field
func handleString(field reflect.Value, value string) (err error) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = strToInt64(field, value)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		err = strToUint64(field, value)

	case reflect.Struct:
		err = strToTime(field, value)
	default:
		err = fmt.Errorf(errors.TypeMismatchWithDetail, field.Kind().String())
	}
	return err
}
