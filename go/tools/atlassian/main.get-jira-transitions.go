package main

import (
	"flag"
	"github.com/sam-caldwell/monorepo/go/ansi"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraTransition"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/version"
)

func main() {
	var app JiraTransition.JiraTransition
	var client Atlassian.Client
	debug := flag.Bool("debug", false, "enable debug output")
	noop := flag.Bool("noop", false, "run command with no effect")
	ver := flag.Bool("version", false, "print version and exit")
	apiKey := flag.String("apiKey", "", "Jira API key")
	issueKey := flag.String("issueKey", "", "Jira issue key")
	domain := flag.String("domain", "", "Jira domain")

	flag.Parse()

	if *ver {
		ansi.Blue().Println(version.Version).Fatal(exit.Success).Reset()
	}

	if issueKey == nil || *issueKey == "" {
		ansi.Red().Println("Issue key cannot be empty or blank.  Use --issueKey").
			Fatal(exit.MissingArg).Reset()
	}

	if apiKey == nil || *apiKey == "" {
		ansi.Red().Println("apiKey cannot be empty or blank.  Use --apiKey").
			Fatal(exit.MissingArg).Reset()
	}

	client.SetDebug(*debug)
	client.SetNoop(*noop)

	if err := client.SetDomain(domain); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	}

	if err := client.SetApiKey(apiKey); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	}

	if err := app.Init(&client, issueKey); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	}

	if transitions, err := app.GetTransitionNames(); err != nil {
		ansi.Red().Println(err.Error()).Fatal(exit.GeneralError).Reset()
	} else {
		ansi.Blue()
		for key, value := range *transitions {
			ansi.Printf("'%s': '%s'", key, value).LF()
		}
		ansi.Reset()
	}
}