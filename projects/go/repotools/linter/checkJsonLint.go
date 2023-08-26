package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
)

func checkJsonLint() error {
	_, response := systemrecon.HasExecutable("jsonlint")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("jsonlint is not installed")
}
