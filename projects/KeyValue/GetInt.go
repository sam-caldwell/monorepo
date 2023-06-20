package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/numbers"
	"reflect"
)

/*
 * KeyValue.GetInt
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as int or typecheck error
 */

// GetInt - Return the value corresponding to a given key as int or typecheck error
func (kv *KeyValue) GetInt(key string) (value int, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		return numbers.IntZero, fmt.Errorf(errors.NotFound)
	}
	if reflect.TypeOf(raw).Kind() != reflect.Int {
		return numbers.IntZero, fmt.Errorf(errors.TypeMismatch)
	}
	return raw.(int), nil
}
