package repobuilder

import repocli "github.com/sam-caldwell/go/v2/projects/repotools/ui"

type BuildRunner func(
	name *string,
	outputDirectory *string,
	notice repocli.NoticeMessagePrintFunc,
	pass repocli.PassMessagePrintFunc,
	fail repocli.FailMessagePrintFunc,
	projectPath *string) error
