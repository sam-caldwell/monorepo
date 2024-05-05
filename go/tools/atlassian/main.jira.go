package main

import (
    "flag"
    "fmt"
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
    "github.com/sam-caldwell/monorepo/go/exit"
    "github.com/sam-caldwell/monorepo/go/simpleArgs"
    "github.com/sam-caldwell/monorepo/go/tools/atlassian/JiraActions"
)

type action func() error

func main() {

	command := simpleArgs.GetCommand("command (create,read,update,delete,list")
	object := simpleArgs.GetCommand("object (issue,project)")

	var fn action

	switch object {
	case "issue":
		fn = func() error {
			var app JiraIssue.Issue

			debug := flag.Bool("debug", false, "enable debug output")
			noop := flag.Bool("noop", false, "run command with no effect")
			apiKey := flag.String("--apiKey", "", "Jira API key")
			domain := flag.String("--domain", "", "Jira domain")
			var descriptor, issueOrKey, jqlString *string
			switch command {
			case "create":
				descriptor = flag.String("--descriptor", "", "Jira issue descriptor file")
			case "update":
				descriptor = flag.String("--descriptor", "", "Jira issue descriptor file")
				issueOrKey = flag.String("--issue", "", "jira issue id or key")
			case "read", "delete":
				issueOrKey = flag.String("--issue", "", "jira issue id or key")
			case "list":
				jqlString = flag.String("--jql", "", "jira jql string")
			default:
				return fmt.Errorf("invalid command")
			}
			flag.Parse()

			if err := app.Init(*debug, *noop, apiKey, domain, descriptor, issueOrKey, jqlString); err != nil {
				return err
			}
			switch command {
			case "create":
				if err := JiraActions.IssueCreate(&app); err != nil {
					return err
				}
			case "read":
				if err := JiraActions.IssueRead(&app); err != nil {
					return err
				}
			case "update":
				if err := JiraActions.IssueUpdate(&app); err != nil {
					return err
				}
			case "delete":
				if err := JiraActions.IssueDelete(&app); err != nil {
					return err
				}
			case "list":
			default:
				return fmt.Errorf("cannot execute invalid command")
			}
			return nil
		}
	case "project":
		fn = func() error {
			return fmt.Errorf("not implemented")
		}
	default:
		ansi.Red().Println("invalid object").Fatal(exit.GeneralError).Reset()
	}
	exit.TerminateOnError(fn())

}
