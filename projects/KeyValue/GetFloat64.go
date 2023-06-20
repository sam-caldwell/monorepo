package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/numbers"
	"reflect"
)

/*
 * KeyValue.GetFloat64
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as float64 or typecheck error
 */

// GetFloat64 - Return the value corresponding to a given key as float64 or typecheck error
func (kv *KeyValue) GetFloat64(key string) (value float64, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		return numbers.Float64zero, fmt.Errorf(errors.NotFound)
	}
	if reflect.TypeOf(raw).Kind() != reflect.Float64 {
		return numbers.Float64zero, fmt.Errorf(errors.TypeMismatch)
	}
	return raw.(float64), nil
}
