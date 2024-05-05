package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// ReadIssue - Read a new jira issue/ticket
func ReadIssue(app *JiraIssue.Issue, issueIdOrKey string) error {

	output, err := app.Read(issueIdOrKey)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
