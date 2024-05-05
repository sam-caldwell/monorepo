package JiraIssue

import (
	"net/http"
)

// Read - read the specified issue
func (jira *Issue) Read(issueIdOrKey string) (output []byte, err error) {

	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	jira.IssueKey = issueIdOrKey

	return jira.client.Send(http.MethodGet, path, jira.Marshall())

}
