package JiraActions

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    "github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// IssueList - List a new jira issue/ticket
func IssueList(app *JiraIssue.Issue) error {

	output, err := app.List()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
