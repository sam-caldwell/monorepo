package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
)

const (
	createCmd = "create"
	readCmd   = "read"
	updateCmd = "update"
	deleteCmd = "delete"
	listCmd   = "list"
)

const (
	issue   = "issue"
	project = "project"
)

// Router - route the execution to the appropriate jiraAction based on command and object
func Router(command *string, object *string) error {
	switch *object {

	case issue:
		return Issue(command, object)

	case project:
		return Project(command, object)

	default:
		ansi.Red().Println("invalid object").Fatal(exit.GeneralError).Reset()

	}
	return nil
}
