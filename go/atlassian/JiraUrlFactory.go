package Atlassian

import (
	"fmt"
	"path"
)

// JiraUrlFactory - form the JIRA URL
func JiraUrlFactory(pattern, tenant, apiPath string) string {

	return path.Join(fmt.Sprintf(pattern, tenant, atlassianFqdn), apiPath)
}
