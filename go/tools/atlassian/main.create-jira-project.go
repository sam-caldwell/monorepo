package main

import (
    "github.com/sam-caldwell/monorepo/go/atlassian/JiraProject"
    "github.com/sam-caldwell/monorepo/go/exit"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/cli"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/commands"
)

func main() {

    const usage = "" +
        "create-jira-project " +
        "--domain <name> " +
        "--descriptor <filename.yaml> " +
        "[--api-key <string>]"

    var app cli.JiraClient[JiraProject.Project]

    cli.GetHelp(usage)
    exit.TerminateOnError(app.GetApiKey())
    exit.TerminateOnError(app.GetDomain())
    exit.TerminateOnError(app.GetDescriptor())
    exit.TerminateOnError(app.Execute(commands.Create))
}
