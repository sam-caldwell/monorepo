package repobuilder

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
)

func BuildCpp(
	name *string,
	outputDirectory *string,
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc,
	projectPath *string) (err error) {

	return fmt.Errorf(errors.NotImplemented)
}
