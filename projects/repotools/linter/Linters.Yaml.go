package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

// Yaml - Lint a given yaml file
func Yaml(filename string) (err error) {
	var out string
	out, err = runcommand.ShellExecute(fmt.Sprintf("yamllint %s", filename))
	fmt.Printf("Lint %s: %s\n", filename, out)
	return err
}
