package repotester

import (
	"fmt"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/go/runcommand"
)

// TestGolang - Run golang test
func TestGolang(
	name *string,
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc,
	projectPath *string) (err error) {

	var stdout string

	command := fmt.Sprintf("go test -failfast -v -count=2 -vet=all %s/...", *projectPath)

	stdout, err = runcommand.ShellExecute(command)
	notice(stdout)
	if err == nil {
		pass(*name, *projectPath)
	} else {
		err = fmt.Errorf("out:%s\nerr:%s", stdout, err)
		fail(*name, *projectPath, err)
	}
	return err
}
