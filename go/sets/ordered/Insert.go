package ordered

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Insert - a simple (unexported insert method)

func (set *Set[T]) Insert(item T, pos int) error {
	set.lock.Lock()
	defer set.lock.Unlock()

	if set.seenBefore(&item) {
		return fmt.Errorf(errors.DuplicateEntry)
	}

	if pos >= len(set.data) || pos < 0 {
		if len(set.data) == 0 && pos == 0 {
			set.data = append(set.data, item)
			return nil
		}
		return fmt.Errorf(errors.IndexOutOfRange)
	}

	// Insert the item at the specified position
	set.data = append(set.data[:pos+1], set.data[pos:]...)
	set.data[pos] = item

	return nil
}
