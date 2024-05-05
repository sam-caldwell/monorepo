package JiraActions

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// Issue - Route issue-related operation
func Issue(command, object *string) error {
	var app JiraIssue.Issue
	var debug, noop *bool
	var apiKey, descriptor, domain, issueOrKey, jqlString *string
	debug = flag.Bool("debug", false, "enable debug output")
	noop = flag.Bool("noop", false, "run command with no effect")
	apiKey = flag.String("apiKey", "", "Jira API key")
	domain = flag.String("domain", "", "Jira domain")
	GetDescriptor := func() *string {
		return flag.String("descriptor", "", "Jira issue descriptor file")
	}
	GetIssueOrKey := func() *string {
		return flag.String("issue", "", "jira issue id or key")
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
		jqlString = flag.String("jql", "", "jira jql string")
	default:
		return fmt.Errorf("invalid command")
	}
	flag.Parse()

	if *debug {
		ansi.Blue().
			Println("debugging...").
			Line("-", 40).
			Printf("       debug: '%v'", *debug).LF().
			Printf("        noop:  '%v'", *noop).LF().
			Printf("     command: '%s'", *command).LF().
			Printf("      object: '%s'", *object).LF().
			Printf("      apiKey: '%s'", *apiKey).LF().
			Printf("      domain: '%s'", *domain).LF()
		if descriptor != nil {
			ansi.Printf("  descriptor: '%s'", *descriptor).LF()
		} else {
			ansi.Printf("  descriptor: ''").LF()
		}
		if issueOrKey != nil {
			ansi.Printf("  issueOrKey: '%s'", *issueOrKey).LF()
		} else {
			ansi.Printf("  issueOrKey: ''").LF()
		}
		if jqlString != nil {
			ansi.Printf("  JsqlString: '%s'", *jqlString).LF()
		} else {
			ansi.Printf("  JsqlString: ''").LF()
		}
		ansi.Line("-", 40).LF().
			Reset()
	}

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
}
