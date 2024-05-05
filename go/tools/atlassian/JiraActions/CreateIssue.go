package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// CreateIssue - Create a new jira issue/ticket
func CreateIssue(app *JiraIssue.Issue) error {

	output, err := app.Create()
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
