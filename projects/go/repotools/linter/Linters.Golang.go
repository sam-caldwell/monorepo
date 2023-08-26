package repolinter

import (
	"fmt"
	monorepo2 "github.com/sam-caldwell/go/v2/projects/go/__system__"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
	"github.com/sam-caldwell/go/v2/projects/go/sets/simple"
	"path/filepath"
)

func GolangLinterFactory() func(filename string) (err error) {

	var directoriesSeenBefore simple.Set

	golang := func(filename string) (err error) {
		var out string
		thisDirectory := filepath.Dir(filename)

		if directoriesSeenBefore.Has(thisDirectory) {
			return nil // don't scan a directory more than once.
		}

		for _, goos := range monorepo2.GetSupportedOpsys() {
			for _, goarch := range monorepo2.GetSupportedArch() {
				command := fmt.Sprintf("GOOS=%s GOARCH=%s staticcheck %s", goos, goarch, filename)
				out, err = runcommand.ShellExecute(command)
				if err != nil {
					return fmt.Errorf("\nerror(goos:%s,goarch:%s):\n\t%s\n\t%v", goos, goarch, out, err)
				}
			}
		}
		return err
	}
	return golang

}
