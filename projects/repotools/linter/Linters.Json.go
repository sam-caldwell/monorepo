package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

func Json(filename string) (err error) {
	var out string
	out, err = runcommand.ShellExecute(fmt.Sprintf("jsonlint --quiet %s", filename))
	if err != nil {
		return err
	}
	if out != "" {
		return fmt.Errorf("\nerror:\n\t'%s'", out)
	}
	return nil
}
