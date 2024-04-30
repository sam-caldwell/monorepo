package cli

import (
	"fmt"
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Delete - Delete a JIRA project or ticket
func (jira *JiraClient) Delete() (err error) {
	switch jira.object {
	case Ticket:
		var descriptor AtlassianJira.Ticket
		return descriptor.Delete(jira.apiKey, jira.GetTicketId())
	case Project:
		var descriptor AtlassianJira.Project
		return descriptor.Delete(jira.apiKey, jira.GetProjectName())
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
