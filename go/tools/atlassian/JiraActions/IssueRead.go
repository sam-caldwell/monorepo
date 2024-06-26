package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
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
