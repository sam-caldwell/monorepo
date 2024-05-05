package JiraIssue

import (
	"net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira *Issue) Create() (output []byte, err error) {

	const path = "/rest/api/2/issue/"

	return jira.client.Send(http.MethodPost, path, jira.Marshall())

}
