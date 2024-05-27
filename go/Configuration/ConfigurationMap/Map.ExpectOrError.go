package configuration

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// ExpectOrPanic - Lookup a given key and return the value or panic.
//
//	 If the given key does not exist or the map is not initialized
//	 the method will panic()
//
//		(c) 2023 Sam Caldwell.  MIT License
func (cfg *Map[K, V]) ExpectOrPanic(name K) V {
	if cfg.data == nil {
		panic(errors.NotInitialized)
	}
	cfg.lock.Lock()
	record, ok := cfg.data[name]
	cfg.lock.Unlock()
	if !ok {
		panic(fmt.Sprintf("missing %s", name))
	}
	return record
}
