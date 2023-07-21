package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func checkJsonLint() error {
	_, response := systemrecon.HasExecutable("jsonlint")
	if response == words.Yes {
		return nil
	}
	return fmt.Errorf("jsonlint is not installed")
}
