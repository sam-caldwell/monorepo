package JiraActions

import "flag"

// getIssueOrKey - Get the Issue or Key (--issue) from cli
func getIssueOrKey() *string {
	return flag.String("issue", "", "jira issue id or key")
}
