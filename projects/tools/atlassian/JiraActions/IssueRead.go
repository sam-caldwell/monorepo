package JiraActions

import (
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/atlassian/JiraIssue"
)

// IssueRead - Read a new jira issue/ticket
func IssueRead(app *JiraIssue.Issue) error {

	output, err := app.Read()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
