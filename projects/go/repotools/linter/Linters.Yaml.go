package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
)

// Yaml - Lint a given yaml file
func Yaml(filename string) (err error) {
	var out string
	out, err = runcommand.ShellExecute(fmt.Sprintf("yamllint %s", filename))
	if err != nil {
		return err
	}
	if out != "" {
		return fmt.Errorf("\nerror:\n"+
			"\t%s\n"+
			"\t%s", out, err)
	}
	return nil
}
