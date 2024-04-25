package persistence

import (
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
)

// QueryCollector - watch a given path
type QueryCollector struct {
	count types.Sequence

	queue map[uuid.UUID][]string //map[clientId]filename

	path directory.Path
}
