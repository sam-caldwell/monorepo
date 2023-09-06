package file

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/fs"
	"os"
)

// CreateTempFile - Create a temporary file in a temporary directory (uses uuid for the context)
func CreateTempFile() (tempFile *os.File, err error) {
	if tempFile, err = os.CreateTemp(os.TempDir(), fmt.Sprintf("t%s-*", uuid.New().String())); err != nil {
		return nil, fmt.Errorf(fs.ErrCouldNotCreateFile, err)
	}
	return tempFile, nil
}
