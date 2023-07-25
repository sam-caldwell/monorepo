package repotester

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
)

func TestNodeJs(
	name *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) (err error) {

	return fmt.Errorf(errors.NotImplemented)
}
