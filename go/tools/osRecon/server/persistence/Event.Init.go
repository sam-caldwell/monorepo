package persistence

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
)

// Init - initialize the event collector object.
func (evt *EventCollector) Init(path *directory.Path) error {
	evt.id = 0
	if !path.Exists() {
		return fmt.Errorf("event collector path must exist (not found)")
	}
	evt.path = *path
	return nil
}
