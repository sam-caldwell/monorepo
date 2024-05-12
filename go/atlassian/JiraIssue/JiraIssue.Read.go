package JiraIssue

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Read - read the specified issue
func (jira *Issue) Read() (output []byte, err error) {

	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	output, err = jira.client.Send(
		http.MethodGet,
		fmt.Sprintf(path, jira.IssueKey),
		jira.Marshall())

	if err != nil {
		return output, err
	}
	if err = jira.Unmarshall(output); err != nil {
		return output, err
	}
	return json.MarshalIndent(jira, "", "  ")
}
