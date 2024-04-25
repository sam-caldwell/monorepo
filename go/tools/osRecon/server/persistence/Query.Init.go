package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
)

// Init - initialize the query collector object.
func (qry *QueryCollector) Init(path *directory.Path) error {

	qry.count = 0

	if !path.Exists() {
		return fmt.Errorf("query collector path must exist (not found)")
	}

	qry.path = *path

	qry.queue = make(map[uuid.UUID][]string)

	return nil

}
