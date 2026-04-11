package simple

import "github.com/sam-caldwell/monorepo/projects/misc"

// Set - Create map of any to its type
type Set[T comparable] struct {
	data map[T]misc.NullObjectStruct
}
