package JiraConfig

import (
	Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"
)

// JiraConfig - Global Jira Configuration
type JiraConfig struct {
	client     *Atlassian.Client
	Properties []Property
}
