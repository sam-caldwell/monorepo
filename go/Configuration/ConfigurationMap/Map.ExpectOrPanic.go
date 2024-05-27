package configuration

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// ExpectOrError - Lookup a given key and return the value and error state.
//
//	 If the given key does not exist or the map is not initialized
//	 the method will return the appropriate error state
//
//		(c) 2023 Sam Caldwell.  MIT License
func (cfg *Map[K, V]) ExpectOrError(name K) (V, error) {
	if cfg.data == nil {
		panic(errors.NotInitialized)
	}
	cfg.lock.Lock()
	record, ok := cfg.data[name]
	cfg.lock.Unlock()
	if !ok {
		var emptyValue V
		return emptyValue, fmt.Errorf("missing %s", name)
	}
	return record, nil
}
