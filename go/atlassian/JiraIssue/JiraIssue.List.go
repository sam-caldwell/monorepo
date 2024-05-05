package JiraIssue

import (
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
)

// List - list issues given a JQL Query object
func (jira *Issue) List(jql *AtlassianTypes.JqlQuery) (output []byte, err error) {
	const path = "/rest/api/2/issue/"

	//ToDo: how do we send the JQL?

	return jira.client.Send(http.MethodPost, path, jira.Marshall())
}
