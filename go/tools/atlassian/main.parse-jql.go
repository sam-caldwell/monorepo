package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"github.com/sam-caldwell/monorepo/go/atlassian/JQL"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/version"
)

func main() {
	var client Atlassian.Client
	var jql JQL.JQL

	debug := flag.Bool("debug", false, "enable debug output")
	noop := flag.Bool("noop", false, "run command with no effect")
	ver := flag.Bool("version", false, "print version and exit")
	apiKey := flag.String("apiKey", "", "Jira API key")
	domain := flag.String("domain", "", "Jira domain")
	username := flag.String("username", "", "Jira user name")
	jqlString := flag.String("jql", "", "jira jql string")

	flag.Parse()

	if *ver {
		ansi.Blue().Println(version.Version).Fatal(exit.Success).Reset()
	}

	if jqlString == nil || *jqlString == "" {
		ansi.Red().Println("jql string is required").Fatal(exit.GeneralError).Reset()
	}

	client.SetDebug(*debug)
	client.SetNoop(*noop)

	exit.TerminateOnError(client.SetUsername(username))
	exit.TerminateOnError(client.SetApiKey(apiKey))
	exit.TerminateOnError(client.SetDomain(domain))
	exit.TerminateOnError(jql.Init(&client))

	if err := jql.Parse(jqlString); err != nil {
		ansi.Red().Printf("Parse Failed: %s", err).LF().Fatal(exit.GeneralError).Reset()
	}
	if output, err := jql.PrettyString(); err != nil {
		ansi.Red().Printf("Failed: %s", err).LF().Fatal(exit.GeneralError).Reset()
	} else {
		ansi.Green().Printf("%s", output).Reset().LF().Reset()
	}
}
