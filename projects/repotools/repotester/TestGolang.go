package repotester

import (
	"fmt"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
	"github.com/sam-caldwell/go/v2/projects/runcommand"
)

// TestGolang - Run golang test
func TestGolang(
	name *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
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
