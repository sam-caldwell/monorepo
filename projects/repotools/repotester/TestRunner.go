package repotester

import repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"

type TestRunner func(
	name *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) error
