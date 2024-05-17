package simple

import (
	"github.com/sam-caldwell/monorepo/go/misc"
)

// Add - add item to set (all items must be of the same type).
func (set *Set[T]) Add(item T) error {
	set.Init()
	set.data[item] = misc.NullObjectStruct{}
	return nil
}
