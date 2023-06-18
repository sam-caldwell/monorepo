package lock

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"os"
	"path/filepath"
	"time"
)

// Create - create a lock file for the contextId
func (lock *File) Create() (contextId string, err error) {

	var handle *os.File = nil
	defer func() {
		if handle != nil {
			_ = handle.Close()
		}
	}()

	for i := 0; i < createRetry; i++ {

		contextId = uuid.New().String()

		fileName := filepath.Join(os.TempDir(), fmt.Sprintf("lock-%s", contextId))
		if file.Exists(fileName) {
			time.Sleep(createRetryWait)
			continue // Try again
		}

		if handle, err = os.CreateTemp(os.TempDir(), fileName); err != nil {
			if handle != nil {
				_ = handle.Close() //make sure we are closed
			}
			time.Sleep(createRetryWait)
			continue //Try again
		} else {
			return contextId, err
		}
	} /*end of for*/

	// Our retries have expired and something terrible is preventing us from

	return "", err
}
