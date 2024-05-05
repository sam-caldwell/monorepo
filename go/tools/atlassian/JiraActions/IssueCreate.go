package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// IssueCreate - Create a new jira issue/ticket
func IssueCreate(app *JiraIssue.Issue) error {

	output, err := app.Create()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
