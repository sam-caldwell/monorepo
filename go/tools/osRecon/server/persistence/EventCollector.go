package persistence

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
)

// EventCollector - Store the event collection to disk
type EventCollector struct {
	id   types.Sequence
	path directory.Path
}
