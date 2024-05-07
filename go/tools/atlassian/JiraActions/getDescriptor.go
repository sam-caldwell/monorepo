package JiraActions

import "flag"

// getDescriptor - get the --descriptor filename from the cli
func getDescriptor() *string {
	return flag.String("descriptor", "", "Jira issue descriptor file")
}
