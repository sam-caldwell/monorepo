package cli

import (
	"fmt"
	AtlassianJira "github.com/sam-caldwell/monorepo/go/atlassian/jira"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Update - Update ticket/project with internal state
func (jira *JiraClient) Update() (err error) {
	switch jira.object {
	case Ticket:
		var descriptor AtlassianJira.Ticket
		if err = descriptor.Load(jira.descriptor); err == nil {
			return descriptor.Update(jira.apiKey)
		}
	case Project:
		var descriptor AtlassianJira.Project
		if err = descriptor.Load(jira.descriptor); err == nil {
			return descriptor.Update(jira.apiKey)
		}
	default:
		err = fmt.Errorf(errors.InvalidCommand)
	}
	return err
}
