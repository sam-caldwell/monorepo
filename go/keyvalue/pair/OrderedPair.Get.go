package pair

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Get - return the pair at the given position or throw an error
//
//	(c) 2023 Sam Caldwell.  MIT License
func (o *OrderedPair[KeyType, ValueType]) Get(pos int) (result Pair[KeyType, ValueType], err error) {
	o.lock.Lock()
	defer o.lock.Unlock()
	endPosition := len(o.data) - 1
	if pos > endPosition || pos < 0 {
		return Pair[KeyType, ValueType]{}, fmt.Errorf(errors.IndexOutOfRange)
	}
	return o.data[pos], nil
}
