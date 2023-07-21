package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"path/filepath"
	"strings"
)

func (linter *Linters) Run(fileName string) error {
	fileExtension := strings.TrimPrefix(filepath.Ext(fileName), words.Period)
	if entry, ok := linter.table[fileExtension]; ok {
		return entry.runner(fileName)
	}
	return fmt.Errorf(noLinter)
}
