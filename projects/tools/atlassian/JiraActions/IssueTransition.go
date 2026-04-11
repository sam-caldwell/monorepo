package JiraActions

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/projects/ansi"
	"github.com/sam-caldwell/monorepo/projects/atlassian/JiraIssue"
)

// IssueTransition - Transition jira issue in the workflow
func IssueTransition(app *JiraIssue.Issue, step *string) error {

	if step == nil || *step == "" {
		return fmt.Errorf("invalid workflow step")
	}

	output, err := app.Transition(step)
	if err != nil {
		return err
	}

	ansi.Green().Println(string(output)).Reset()

	return nil
}
