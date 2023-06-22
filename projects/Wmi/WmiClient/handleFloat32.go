//go:build windows
// +build windows

package wmiclient

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

// handleFloat32 cast a float32 value into the field
func handleFloat32(field reflect.Value, value float32) error {
	if field.Kind() == reflect.Float32 {
		field.SetFloat(float64(value))
		return nil
	}
	return fmt.Errorf(errors.TypeMismatchWithDetail, reflect.Float32.String())
}
