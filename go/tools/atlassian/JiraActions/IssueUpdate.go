package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// IssueUpdate - Update a new jira issue/ticket
func IssueUpdate(app *JiraIssue.Issue) error {

	output, err := app.Update()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
