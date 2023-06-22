//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

// handleFloat64 cast a float64 value into the field
func handleFloat64(field reflect.Value, value float64) error {
	if field.Kind() == reflect.Float64 {
		field.SetFloat(value)
		return nil
	}
	return fmt.Errorf(errors.TypeMismatchWithDetail, reflect.Float64.String())
}
