package repolinter

import (
	"fmt"
	"path/filepath"
)

func (linter *Linters) Run(fileName string) error {
	fileExtension := filepath.Ext(fileName)
	if fn, ok := linter.table[fileExtension]; ok {
		return fn()
	}
	return fmt.Errorf(noLinter)
}
