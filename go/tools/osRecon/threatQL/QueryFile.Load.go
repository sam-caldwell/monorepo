package threatQL

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/types"
)

// Load - Load the first query file in the given path and clientId queue subdirectory
// if a query file is already loaded,
//
// Note that if err != nil, the result may be nil.
func (q *QueryFile) Load(path directory.Path, clientId types.ClientId) (result *Query, err error) {
	const (
		maxQueryFileSz = 16384
	)
	if err := q.Prune(); err != nil {
		return nil, err
	}

	if err = q.getNext(path, clientId); err != nil {
		return nil, err
	}

	if err = q.readFile(); err != nil {
		return nil, err
	}
	return q, nil
}
