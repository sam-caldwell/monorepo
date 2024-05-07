package AtlassianTypes

import (
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

func (jira *JiraTransition) Init(client *Atlassian.Client) error {
	if client == nil {
		return fmt.Errorf("client is nil")
	}
	jira.client = client
}
