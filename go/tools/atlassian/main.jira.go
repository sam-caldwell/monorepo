package main

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/exit"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/cli"
)

func main() {
    ansi.Blue().Println("Starting...").Reset()
    var app cli.JiraClient

    app.GetHelp(cli.JiraClientUsage)

    exit.TerminateOnError(app.Configure())

    exit.TerminateOnError(app.Execute())
}
