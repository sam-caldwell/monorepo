package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
)

/*
 * keyvalue.GetBool
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as boolean (if appropriate) or typecheck error
 */

// GetBool - Return the value corresponding to a given key as boolean (if appropriate) or typecheck error
func (kv *KeyValue) GetBool(key string) (value bool, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		return false, fmt.Errorf(errors.NotFound)
	}
	if reflect.TypeOf(raw).Kind() != reflect.Bool {
		return false, fmt.Errorf(errors.TypeMismatch)
	}
	return raw.(bool), nil
}
