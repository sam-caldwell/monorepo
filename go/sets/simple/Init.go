package simple

import "github.com/sam-caldwell/monorepo/go/misc"

// Init - Initialize the state's internal data structure
func (set *Set[T]) Init() {
	// If our set is not initialized, let's do it now.
	// We do this here to avoid nil exceptions EVER (we prefer safety to speed)
	if set.data == nil {
		set.data = make(map[T]misc.NullObjectStruct)
	}
}
