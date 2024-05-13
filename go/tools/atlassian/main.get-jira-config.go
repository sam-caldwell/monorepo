package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraConfig"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/version"
)

func main() {
	var client Atlassian.Client
	var config JiraConfig.JiraConfig

	debug := flag.Bool("debug", false, "enable debug output")
	noop := flag.Bool("noop", false, "run command with no effect")
	ver := flag.Bool("version", false, "print version and exit")
	apiKey := flag.String("apiKey", "", "Jira API key")
	domain := flag.String("domain", "", "Jira domain")
	username := flag.String("username", "", "Jira user name")

	flag.Parse()

	if *ver {
		ansi.Blue().Println(version.Version).Fatal(exit.Success).Reset()
	}

	client.SetDebug(*debug)
	client.SetNoop(*noop)

	exit.TerminateOnError(client.SetUsername(username))
	exit.TerminateOnError(client.SetApiKey(apiKey))
	exit.TerminateOnError(client.SetDomain(domain))
	exit.TerminateOnError(config.Init(&client))
	ansi.Green().Println(config.String()).LF().Reset()
}
