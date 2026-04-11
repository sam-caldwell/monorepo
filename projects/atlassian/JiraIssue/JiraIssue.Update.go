package JiraIssue

import (
	"net/http"
)

// Update -  Update the identified issue as defined by the internal Issue struct state
func (jira *Issue) Update() (output []byte, err error) {
	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	return jira.client.Send(http.MethodPut, path, jira.Marshall())
}
