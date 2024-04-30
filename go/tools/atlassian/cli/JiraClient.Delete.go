package cli

import (
	"fmt"
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

// Delete - Delete a JIRA project or ticket
func (jira *JiraClient) Delete() (err error) {
	switch jira.object {
	case Ticket:
		ticketId, err := simpleArgs.GetOptionValue("--ticket")
		if err != nil {
			return err
		}
		var descriptor AtlassianJira.Ticket
		return descriptor.Delete(jira.apiKey, AtlassianJira.TicketId(ticketId))
	case Project:
		projectName, err := simpleArgs.GetOptionValue("--project")
		if err != nil {
			return err
		}
		var descriptor AtlassianJira.Project
		return descriptor.Delete(jira.apiKey, projectName)
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
