package lock

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/fs/file"
	os2 "github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"path/filepath"
)

// Free - delete the lock file for contextId
func (lock *File) Free(contextId string) error {
	if f := filepath.Join(os2.TempDir(), fmt.Sprintf("lock-%s", contextId)); file.Exists(f) {
		return os2.Remove(f)
	}
	return nil
}
