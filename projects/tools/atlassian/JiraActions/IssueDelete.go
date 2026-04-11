package JiraActions

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/atlassian/JiraIssue"
)

// IssueDelete - Delete a new jira issue/ticket
func IssueDelete(app *JiraIssue.Issue) error {

	output, err := app.Delete()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
