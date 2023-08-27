package repobuilder

import (
	repocli2 "github.com/sam-caldwell/go/v2/projects/go/repotools/ui"
)

type BuildRunner func(
	name *string,
	outputDirectory *string,
	notice repocli2.NoticeMessagePrintFunc,
	pass repocli2.PassMessagePrintFunc,
	fail repocli2.FailMessagePrintFunc,
	projectPath *string) error
