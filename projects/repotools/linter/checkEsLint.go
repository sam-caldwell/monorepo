package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkEsLint() error {
	_, response := systemrecon.HasExecutable("npx eslint -v")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("eslint is not installed")
}
