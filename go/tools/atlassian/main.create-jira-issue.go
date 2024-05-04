package main

import (
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/atlassian/cli"
	"github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
)

func main() {

	const usage = "" +
		"create-jira-ticket " +
		"--domain <name> " +
		"--descriptor <filename.yaml> " +
		"[--api-key <string>]"

	var app cli.JiraClient[JiraIssue.Issue]

	cli.GetHelp(usage)
	app.SetDebug(cli.GetDebugFlag())
	app.SetNoop(cli.GetNoopFlag())
	exit.TerminateOnError(app.GetApiKey())
	exit.TerminateOnError(app.GetDescriptor())
	exit.TerminateOnError(app.GetDomain())
	exit.TerminateOnError(app.Execute(commands.Create))

}
