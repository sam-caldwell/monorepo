package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkYamlLint() error {
	_, response := systemrecon.HasExecutable("yamllint")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("yamllint is not installed")
}
