package JiraIssue

import (
	"net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira *Issue) Create() (output []byte, err error) {

	const path = "/rest/api/3/issue"

	return jira.client.Send(http.MethodPost, path, jira.Marshall())

	/*
	   expected output:

	   {
	       "id": "11026",
	       "key": "CLITEST-2",
	       "self": "https://samcaldwell.atlassian.net/rest/api/3/issue/11026"
	   }
	*/
}
