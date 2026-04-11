package JiraConfig

import Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"

// Init - initialize the JIRA config object
func (jira *JiraConfig) Init(client *Atlassian.Client) error {
	jira.client = client
	return jira.GetFromServer()
}
