package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
)

/*
 * KeyValue.GetString
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Return the value corresponding to a given key as string.
 */

// GetString - Return the value corresponding to a given key as string.
func (kv *KeyValue) GetString(key string) (value string, err error) {
	raw, ok := kv.FindKey(key)
	if !ok {
		return words.EmptyString, fmt.Errorf(errors.NotFound)
	}
	return fmt.Sprintf("%v", raw), nil
}
