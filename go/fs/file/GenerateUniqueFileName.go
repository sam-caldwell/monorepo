package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"path/filepath"
	"time"
)

// GenerateUniqueFileName - Create a unique filename based on unix timestamp (nanosecond)
func GenerateUniqueFileName(path, extension string, retries uint) (fileName string, err error) {
	for retry := 0; retry < int(retries); retry++ {
		seq := time.Now().UnixNano()
		fileName = filepath.Join(
			path,
			fmt.Sprintf("%d.%s", seq, extension))
		if !Exists(fileName) {
			break // retry only if file exists and we have a collision.
		}
		time.Sleep(1 * time.Nanosecond)
	}
	if Exists(fileName) {
		return words.EmptyString, fmt.Errorf("unrecoverable filename collision")
	}
	return fileName, err
}
