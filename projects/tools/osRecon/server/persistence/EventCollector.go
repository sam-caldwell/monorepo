package persistence

import (
	"github.com/sam-caldwell/monorepo/projects/fs/directory"
	"github.com/sam-caldwell/monorepo/projects/types"
)

// EventCollector - Store the event collection to disk
type EventCollector struct {
	id   types.Sequence
	path directory.Path
}
