package repobuilder

import (
	"fmt"
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
	"path/filepath"
)

func BuildGolang(
	name *string,
	outputDirectory *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) (err error) {

	var stdout string
	for _, goos := range monorepo.GetSupportedOpsys() {
		var extension string
		if goos == "windows" {
			extension = ".exe"
		} else {
			extension = ""
		}
		for _, goarch := range monorepo.GetSupportedArch() {
			var source = filepath.Join(*projectPath, "main.go")

			var binary = filepath.Join(*outputDirectory, goos, goarch, *name+extension)

			command := fmt.Sprintf("go build -o %s %s", binary, source)

			stdout, err = runcommand.ShellExecute(command)
			notice(stdout)
			if err != nil {
				err = fmt.Errorf("out:%s\nerr:%s", stdout, err)
				fail(*name, *projectPath, err)
				return err
			}
			pass(*name, binary)
		}
	}
	return err
}
