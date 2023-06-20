package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/numbers"
	"reflect"
)

/*
 * KeyValue.GetFloat32
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as float or typecheck error
 */

// GetFloat32 - Return the value corresponding to a given key as float or typecheck error
func (kv *KeyValue) GetFloat32(key string) (value float32, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		return numbers.Float32zero, fmt.Errorf(errors.NotFound)
	}
	if reflect.TypeOf(raw).Kind() != reflect.Float32 {
		return numbers.Float32zero, fmt.Errorf(errors.TypeMismatch)
	}
	return raw.(float32), nil
}
