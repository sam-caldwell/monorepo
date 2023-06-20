package ordered

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

func (set *Set) Delete(pos int) error {
	if (pos >= 0) && (pos < len(set.data)) {
		set.lock.Lock()
		defer set.lock.Unlock()
		set.data = append(set.data[:pos], set.data[pos+1:]...)
		return nil
	}
	return fmt.Errorf(errors.IndexOutOfRange)
}
