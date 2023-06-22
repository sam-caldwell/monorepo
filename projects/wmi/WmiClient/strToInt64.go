package wmiclient

import (
	"reflect"
	"strconv"
)

// strToInt64 Cast a string value to int64 and store to the field
func strToInt64(field reflect.Value, value string) error {
	if i, err := strconv.ParseInt(value, 10, 64); err != nil {
		return err
	} else {
		field.SetInt(i)
		return nil
	}
}
