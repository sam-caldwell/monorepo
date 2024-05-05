package JiraActions

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
	"github.com/sam-caldwell/monorepo/go/exit"
)

const (
	createCmd = "create"
	readCmd   = "read"
	updateCmd = "update"
	deleteCmd = "delete"
	listCmd   = "list"
)

const (
	issue   = "issue"
	project = "project"
)

// Router - route the execution to the appropriate jiraAction based on command and object
func Router(command *string, object *string) error {
	switch *object {
	case issue:
		var app JiraIssue.Issue

		debug := flag.Bool("debug", false, "enable debug output")
		noop := flag.Bool("noop", false, "run command with no effect")
		apiKey := flag.String("--apiKey", "", "Jira API key")
		domain := flag.String("--domain", "", "Jira domain")
		var descriptor, issueOrKey, jqlString *string

		GetDescriptor := func() *string {
			return flag.String("--descriptor", "", "Jira issue descriptor file")
		}
		GetIssueOrKey := func() *string {
			return flag.String("--issue", "", "jira issue id or key")
		}
		switch *command {
		case createCmd:
			descriptor = GetDescriptor()
		case updateCmd:
			descriptor = GetDescriptor()
			issueOrKey = GetIssueOrKey()
		case readCmd, deleteCmd:
			issueOrKey = GetIssueOrKey()
		case listCmd:
			jqlString = flag.String("--jql", "", "jira jql string")
		default:
			return fmt.Errorf("invalid command")
		}
		flag.Parse()

		if err := app.Init(*debug, *noop, apiKey, domain, descriptor, issueOrKey, jqlString); err != nil {
			return err
		}

		switch *command {
		case createCmd:
			if err := IssueCreate(&app); err != nil {
				return err
			}
		case readCmd:
			if err := IssueRead(&app); err != nil {
				return err
			}
		case updateCmd:
			if err := IssueUpdate(&app); err != nil {
				return err
			}
		case deleteCmd:
			if err := IssueDelete(&app); err != nil {
				return err
			}
		case listCmd:
			if err := IssueList(&app); err != nil {
				return err
			}
		default:
			return fmt.Errorf("cannot execute invalid command")
		}

	case project:
		return fmt.Errorf("not implemented")

	default:
		ansi.Red().Println("invalid object").Fatal(exit.GeneralError).Reset()
	}
	return nil
}
