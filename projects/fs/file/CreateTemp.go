package file

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/go/v2/projects/fs"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
)

// CreateTempFile - Create a temporary file in a temporary directory (uses uuid for the context)
func CreateTempFile() (string, error) {
	tempFile, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("t%s-*", uuid.New().String()))
	if err != nil {
		return words.EmptyString, fmt.Errorf(fs.ErrCouldNotCreateFile, err)
	}
	defer func() { _ = tempFile.Close() }()
	return tempFile.Name(), nil
}
