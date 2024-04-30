package cli

import (
	"github.com/sam-caldwell/monorepo/go/list"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
	"os"
)

func (jira *JiraClient) Configure() (err error) {
	cli.GetHelp(jiraClientUsage)
	if err = jira.command.FromString(simpleArgs.GetCommand(helpText)); err != nil {
		return err
	}
	os.Args = list.DeleteElement(os.Args, 1)
	if err = jira.object.FromString(simpleArgs.GetCommand(helpText)); err != nil {

	}
	os.Args = list.DeleteElement(os.Args, 1)
	return err
}
