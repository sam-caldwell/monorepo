package lock

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/monorepo/v2/projects/go/fs/file"
	"github.com/sam-caldwell/monorepo/v2/projects/go/wrappers/os"
	"path/filepath"
)

// Check - check to see if the lock file for contextId exists
func (lock *File) Check(contextId string) error {
	if file.Exists(filepath.Join(os.TempDir(), fmt.Sprintf("lock-%s", contextId))) {
		return fmt.Errorf(errors.LockCheckFailed)
	}
	return nil
}
