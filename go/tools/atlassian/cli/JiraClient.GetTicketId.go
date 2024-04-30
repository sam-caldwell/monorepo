package cli

import (
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

func (jira *JiraClient) GetTicketId() AtlassianJira.TicketId {
	value, err := simpleArgs.GetOptionValue("--ticket")

	if err != nil {
		if err.Error() != simpleArgs.OptionNotFound {
			value = AtlassianJira.DefaultJqlQuery
		}
	}
	return AtlassianJira.TicketId(value)
}
