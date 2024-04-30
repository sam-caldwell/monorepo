package cli

import (
	"fmt"
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

func (jira *JiraClient) Read() (err error) {
	switch jira.object {
	case Ticket:
		var ticket AtlassianJira.Ticket
		return ticket.Read(jira.apiKey, jira.GetTicketId())
	case Project:
		var project AtlassianJira.Project
		return project.Read(jira.apiKey, jira.GetProjectName())
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
