package JiraActions

import "flag"

func getJQLQuery() *string {
	return flag.String("jql", "", "jira jql string")
}
