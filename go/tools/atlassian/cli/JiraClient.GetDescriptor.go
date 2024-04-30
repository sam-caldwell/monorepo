package cli

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// GetDescriptor - Get the --descriptor value if appropriate
func (jira *JiraClient) GetDescriptor() {
	if jira.command == Create || jira.command == Update {
		if thisDescriptor, err := simpleArgs.GetOptionValue("--descriptor"); err != nil {
			ansi.Red().Println(errors.MissingArguments).Fatal(exit.MissingArg).Reset()
		} else {
			if err := jira.descriptor.Set(thisDescriptor); err != nil {
				ansi.Red().Println(errors.MissingArguments).Fatal(exit.MissingArg).Reset()
			}
		}
	}
}
