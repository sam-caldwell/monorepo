package lock

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"os"
	"path/filepath"
)

// Check - check to see if the lock file for contextId exists
func (lock *File) Check(contextId string) error {
	if file.Exists(filepath.Join(os.TempDir(), fmt.Sprintf("lock-%s", contextId))) {
		return fmt.Errorf(exit.ErrLockCheckFailed)
	}
	return nil
}
