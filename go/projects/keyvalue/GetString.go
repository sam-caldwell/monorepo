package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/exit/errors"
	"github.com/sam-caldwell/monorepo/go/projects/misc/words"
)

/*
 * keyvalue.GetString
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
