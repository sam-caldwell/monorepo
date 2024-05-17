package ordered

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Peek - Peek into an element at a given position and return the same
func (set *Set[T]) Peek(pos int) (out T, err error) {
	if pos < 0 || pos >= len(set.data) {
		return out, fmt.Errorf(errors.IndexOutOfRange)
	}
	return set.data[pos], nil
}
