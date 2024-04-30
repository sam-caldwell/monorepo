package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

func (jira *JiraClient) Read() (err error) {
	var sdk Atlassian.Jira

	switch jira.object {
	case Ticket:
		return sdk.ReadTicket(jira.GetJqlQuery())
	case Project:
		return sdk.ReadProject(jira.GetProjectName())
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
