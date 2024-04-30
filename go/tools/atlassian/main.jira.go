package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/tools/atlassian/cli"
)

/*
jira create {ticket,project} <name> --descriptor <filename.yaml> [--api-key <string>]

jira read {ticket, project} <name> [--api-key <string>]

jira update {ticket,project  --descriptor <filename.yaml>

jira delete {ticket, project} <name> [--api-key <string>]

jira list ticket <jql> [--api-key <string>]

jira list projects --offset <int> --count <int> [--api-key <string>]
*/
func main() {
	ansi.Blue().Println("Starting...").Reset()
	var app cli.JiraClient
	exit.TerminateOnError(app.Configure())
}
