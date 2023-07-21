package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"path/filepath"
)

func (linter *Linters) Run(fileName string) error {
	fileExtension := filepath.Ext(fileName)
	if fn, ok := linter.table[fileExtension]; ok {
		return fn()
	}
	return fmt.Errorf(errors.NotImplemented)
}
