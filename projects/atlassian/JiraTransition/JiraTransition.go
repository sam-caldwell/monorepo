package JiraTransition

import Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"

// JiraTransition - The top level object for interacting with Jira Transitions
type JiraTransition struct {
	client      *Atlassian.Client
	transitions Transitions
}
