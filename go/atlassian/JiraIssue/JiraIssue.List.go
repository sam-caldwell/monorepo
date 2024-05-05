package JiraIssue

import (
	"net/http"
)

// List - list issues given a JQL Query object
func (jira *Issue) List() (output []byte, err error) {
	const path = "/rest/api/2/issue/"

	//ToDo: how do we send the JQL?

	return jira.client.Send(http.MethodPost, path, jira.Marshall())
}
