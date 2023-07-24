package repobuilder

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"
)

func BuildC(
	name *string,
	outputDirectory *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) error {

	return fmt.Errorf(errors.NotImplemented)
}
