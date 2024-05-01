package cli

import (
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

func (jira *JiraClient) GetProjectName() (project string) {
	value, err := simpleArgs.GetOptionValue("--project")

	if err != nil {
		if err.Error() != simpleArgs.OptionNotFound {
			value = AtlassianJira.DefaultJqlQuery
		}
	}
	return value
}