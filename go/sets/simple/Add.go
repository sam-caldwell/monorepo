package simple

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/misc"
)

// Add - add item to set (all items must be of the same type).
func (set *Set) Add(item any) error {
	set.Init()
	// Joke: if an item is none, does anyone hear us ignore it in the set?
	//       It's nothing, right?  So... *poof* ...it goes away...
	if item == nil {
		return nil
	}
	// A set must have all items of the same type as the first.
	if !set.TypeCheck(item) {
		return fmt.Errorf(errors.TypeMismatch)
	}

	// Add the item to the set's data
	set.data[item] = misc.NullObjectStruct{}
	return nil
}
