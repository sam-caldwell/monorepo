package configuration

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// ExpectOrPanic - Lookup a given key and return the value or panic.
//
//	If the given key does not exist or the map is not initialized
//	the method will panic()
//
//	(c) 2023 Sam Caldwell.  MIT License.
func (cfg *Map[K, V]) ExpectOrPanic(name K) V {
	if cfg == nil {
		panic(errors.NotInitialized)
	}
	record, ok := (*cfg)[name]
	if !ok {
		panic(fmt.Sprintf("missing %s", name))
	}
	return record
}
