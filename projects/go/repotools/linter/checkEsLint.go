package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
)

func checkEsLint() error {
	_, response := systemrecon.HasExecutable("npx eslint -v")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("eslint is not installed")
}
