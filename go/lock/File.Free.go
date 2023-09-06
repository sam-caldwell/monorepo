package lock

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
	"path/filepath"
)

// Free - delete the lock file for contextId
func (lock *File) Free(contextId string) error {
	if f := filepath.Join(os.TempDir(), fmt.Sprintf("lock-%s", contextId)); file.Exists(f) {
		return os.Remove(f)
	}
	return nil
}
