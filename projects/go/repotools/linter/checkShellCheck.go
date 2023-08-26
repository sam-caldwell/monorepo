package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
)

func checkShellCheck() error {
	_, response := systemrecon.HasExecutable("shellcheck")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("shellcheck is not installed")
}
