package repobuilder

import (
	"fmt"
	monorepo2 "github.com/sam-caldwell/go/v2/projects/go/__system__"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
	"path/filepath"
)

func BuildGolang(
	name *string,
	outputDirectory *string,
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc,
	projectPath *string) (err error) {

	var stdout string
	for _, goos := range monorepo2.GetSupportedOpsys() {
		var extension string
		if goos == "windows" {
			extension = ".exe"
		} else {
			extension = ""
		}
		for _, goarch := range monorepo2.GetSupportedArch() {
			var source = filepath.Join(*projectPath, "main.cpp")

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
