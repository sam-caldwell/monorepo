package lock

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

// Create - create a lock file for the contextId
func (lock *File) Create() (contextId string, err error) {
	const GenerateTempDir = ""

	var handle *os.File = nil
	defer func() {
		if handle != nil {
			_ = handle.Close()
		}
	}()

	for i := 0; i < createRetry; i++ {
		contextId = uuid.New().String()
		handle, err = os.CreateTemp(GenerateTempDir, fmt.Sprintf("lock-%s", contextId))
		if err != nil {
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
