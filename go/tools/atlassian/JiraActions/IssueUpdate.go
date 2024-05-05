package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// UpdateIssue - Update a new jira issue/ticket
func UpdateIssue(app *JiraIssue.Issue, issueIdOrKey string) error {

	output, err := app.Update(issueIdOrKey)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
