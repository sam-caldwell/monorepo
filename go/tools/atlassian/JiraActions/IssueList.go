package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// IssueList - List jira tickets matching the query parameters
//
//	jqlString  - The Jira Query Language string (passed by reference)
//	expand     - a list of additional information fields to expand the result set
//	fields     - a list of fields to return in the result set
//	maxResults - the maximum number of records to return
//	startAt    - the offset from the first matching record to return
func IssueList(
	app *JiraIssue.Issue,
	jqlString *string,
	expand,
	fields *[]string,
	maxResults,
	startAt uint) error {

	if output, err := app.List(jqlString, expand, fields, maxResults, startAt); err != nil {
		return err
	} else {
		ansi.Green().Println(string(output)).Reset()
	}
	return nil
}
