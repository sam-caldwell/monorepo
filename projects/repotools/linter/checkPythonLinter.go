package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkPythonLinter() error {
	_, response := systemrecon.HasExecutable("flake8")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("flake8 is not installed")
}
