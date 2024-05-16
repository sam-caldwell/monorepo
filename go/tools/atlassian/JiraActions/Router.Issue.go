package JiraActions

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/list"
	"github.com/sam-caldwell/monorepo/go/misc"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// Issue - Route issue-related operation
func Issue(command, object *string) error {
	var app JiraIssue.Issue
	var debug, noop *bool
	var apiKey, descriptor, domain, issueOrKey, jqlString, username,
		workflowStep, expand, fields *string
	var maxResults, startAt *uint

	debug = flag.Bool("debug", false, "enable debug output")
	noop = flag.Bool("noop", false, "run command with no effect")
	username = flag.String("username", "", "Jira user name")
	apiKey = flag.String("apiKey", "", "Jira API key")
	domain = flag.String("domain", "", "Jira domain")

	switch *command {
	case createCmd:
		descriptor = getDescriptor()
	case updateCmd:
		descriptor = getDescriptor()
		issueOrKey = getIssueOrKey()
	case readCmd, deleteCmd:
		issueOrKey = getIssueOrKey()
	case listCmd:
		fields = getFieldsList()
		expand = getExpand()
		maxResults = getMaxResults()
		startAt = getStartAt()
		jqlString = getJQLQuery()
	case transitionCmd:
		issueOrKey = getIssueOrKey()
		workflowStep = getWorkflowStep()
	default:
		return fmt.Errorf("invalid command")
	}
	flag.Parse()

	if *debug {
		ansi.Blue().
			Println("debugging...").
			Line("-", 40).
			Printf("       debug: '%v'", *debug).LF().
			Printf("        noop: '%v'", *noop).LF().
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
			ansi.Printf("  JqlString: '%s'", *jqlString).LF()
		} else {
			ansi.Printf("  JqlString: ''").LF()
		}
		ansi.Line("-", 40).LF().
			Reset()
	}

	if err := app.Init(*debug, *noop, username, apiKey, domain, descriptor, issueOrKey); err != nil {
		return err
	}

	switch *command {
	case createCmd:
		if err := requireDescriptorFile(descriptor); err != nil {
			return err
		}
		if err := IssueCreate(&app); err != nil {
			return err
		}
	case readCmd:
		if err := IssueRead(&app); err != nil {
			return err
		}
	case updateCmd:
		if err := requireDescriptorFile(descriptor); err != nil {
			return err
		}
		if err := IssueUpdate(&app); err != nil {
			return err
		}
	case deleteCmd:
		if err := IssueDelete(&app); err != nil {
			return err
		}
	case listCmd:
		if err := exit.ExpectNonEmptyStringP(jqlString); err != nil {
			return err
		}

		expandFields := list.FromStringP(expand, words.Comma, true)

		fieldList := list.FromStringP(fields, words.Comma, true)

		err := IssueList(&app, jqlString, &expandFields, &fieldList,
			misc.PreventNilDereference(maxResults),
			misc.PreventNilDereference(startAt))

		if err != nil {
			return err
		}

	case transitionCmd:
		if err := IssueTransition(&app, workflowStep); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot execute invalid command")
	}
	return nil
}
