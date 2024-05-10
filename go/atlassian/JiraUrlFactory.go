package Atlassian

import (
	"fmt"
)

// JiraUrlFactory - form the JIRA URL
func JiraUrlFactory(pattern, tenant, apiPath string) string {

	return fmt.Sprintf(pattern, tenant, atlassianFqdn, apiPath)
}
