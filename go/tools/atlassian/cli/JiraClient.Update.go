package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

func (jira *JiraClient) Update() (err error) {
	switch jira.object {
	case Ticket:
		return jira.UpdateTicket()
	case Project:
		return jira.UpdateProject()
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
