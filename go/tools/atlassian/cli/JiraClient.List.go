package cli

import (
	"fmt"
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// List - list projects/tickets from JIRA
func (jira *JiraClient) List() (err error) {
	switch jira.object {
	case Ticket:
		var ticket AtlassianJira.Ticket
		return ticket.List(jira.apiKey, jira.GetJqlQuery())
	case Project:
		var project AtlassianJira.Project
		return project.List(jira.apiKey)
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
