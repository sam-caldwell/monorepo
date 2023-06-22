package wmiclient

import (
	"reflect"
	"strconv"
)

// strToUint64 Cast a string value to uint64 and store to the field
func strToUint64(field reflect.Value, value string) error {
	if i, err := strconv.ParseUint(value, 10, 64); err != nil {
		return err
	} else {
		field.SetUint(i)
		return nil
	}
}
