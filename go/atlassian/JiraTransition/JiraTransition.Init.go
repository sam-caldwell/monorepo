package JiraTransition

import (
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

// Init - Given an Atlassian Client and issue key, initialise the internal state of the object
func (jira *JiraTransition) Init(client *Atlassian.Client, issueKey *string) error {
	if client == nil {
		return fmt.Errorf("client is nil")
	}
	jira.client = client

	if err := jira.GetTransitions(issueKey); err != nil {
		return err
	}

	return nil
}
