package repotester

import (
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
)

type TestRunner func(
	name *string,
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc,
	projectPath *string) error
