package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
)

// DeleteIssue - Delete a new jira issue/ticket
func DeleteIssue(app *JiraIssue.Issue, issueIdOrKey string) error {

	output, err := app.Delete(issueIdOrKey)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
