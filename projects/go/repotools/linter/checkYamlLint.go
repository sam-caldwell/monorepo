package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
)

func checkYamlLint() error {
	_, response := systemrecon.HasExecutable("yamllint")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("yamllint is not installed")
}
