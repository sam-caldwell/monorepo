package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// IssueList - List a new jira issue/ticket
func IssueList(app *JiraIssue.Issue, jqlString *string,
	expand, fields *[]string, maxResults, startAt uint) error {

	output, err := app.List(jqlString, expand, fields, maxResults, startAt)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
