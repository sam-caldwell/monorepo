package JiraConfig

import Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"

// Init - initialize the JIRA config object
func (config *JiraConfig) Init(client *Atlassian.Client) error {
	config.client = client
	return config.GetFromServer()
}
