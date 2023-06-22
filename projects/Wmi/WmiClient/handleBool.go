//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

// handleBool cast a bool value into the field
func handleBool(field reflect.Value, value bool) error {
	if field.Kind() == reflect.Bool {
		field.SetBool(value)
		return nil
	}
	return fmt.Errorf(errors.TypeMismatchWithDetail, reflect.Bool.String())
}
