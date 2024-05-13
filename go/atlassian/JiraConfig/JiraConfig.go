package JiraConfig

import (
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

// JiraConfig - Global Jira Configuration
type JiraConfig struct {
	client           *Atlassian.Client
	AdvancedSettings map[string]string
	GlobalSettings   map[string]string
	err              error
}
