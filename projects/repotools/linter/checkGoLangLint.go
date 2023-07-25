package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkGoLangLint() error {
	_, response := systemrecon.HasExecutable("staticcheck")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("staticcheck is not installed")
}
