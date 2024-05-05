package JiraIssue

import (
	"net/http"
)

// Delete the specified issue
func (jira *Issue) Delete() (output []byte, err error) {
	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	return jira.client.Send(http.MethodDelete, path, jira.Marshall())
}
