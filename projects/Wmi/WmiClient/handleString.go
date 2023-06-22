//go:build windows
// +build windows

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
		return strToInt64(field, value)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strToUint64(field, value)

	case reflect.Struct:
		return strToTime(field, value)
	}
	return fmt.Errorf(errors.TypeMismatchWithDetail, reflect.Bool.String())
}
