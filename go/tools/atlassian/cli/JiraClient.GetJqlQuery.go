package cli

import (
	"github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

func (jira *JiraClient) GetJqlQuery() (query AtlassianJira.JQL) {

	rawQuery, err := simpleArgs.GetOptionValue("--jql")

	if err != nil {
		if err.Error() != simpleArgs.OptionNotFound {
			rawQuery = AtlassianJira.DefaultJqlQuery
		}
	}
	return AtlassianJira.JQL(rawQuery)
}
