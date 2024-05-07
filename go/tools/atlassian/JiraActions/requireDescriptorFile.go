package JiraActions

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// requireDescriptorFile - throw an error if the descriptor file is not found.
func requireDescriptorFile(fileName *string) (err error) {
	if !file.Existsp(fileName) {
		err = fmt.Errorf("missing descriptor file")
	}
	return err
}
