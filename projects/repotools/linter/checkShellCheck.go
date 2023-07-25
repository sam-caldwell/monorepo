package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkShellCheck() error {
	_, response := systemrecon.HasExecutable("shellcheck")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("shellcheck is not installed")
}
