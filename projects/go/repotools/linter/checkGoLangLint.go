package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
)

func checkGoLangLint() error {
	_, response := systemrecon.HasExecutable("staticcheck")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("staticcheck is not installed")
}
