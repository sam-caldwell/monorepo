package lock

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"os"
	"path/filepath"
)

// Free - delete the lock file for contextId
func (lock *File) Free(contextId string) error {
	if f := filepath.Join(os.TempDir(), fmt.Sprintf("lock-%s", contextId)); file.Exists(f) {
		return os.Remove(f)
	}
	return nil
}
