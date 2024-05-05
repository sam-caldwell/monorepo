package JiraActions

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/atlassian/JiraIssue"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
)

// IssueList - List a new jira issue/ticket
func IssueList(app *JiraIssue.Issue, jql *AtlassianTypes.JqlQuery) error {

	output, err := app.List(jql)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
