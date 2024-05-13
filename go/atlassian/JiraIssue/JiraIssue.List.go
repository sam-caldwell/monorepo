package JiraIssue

import (
    "github.com/sam-caldwell/monorepo/go/atlassian/JQL"
    "net/http"
)

// List - list issues given a JQL Query object
func (jira *Issue) List(query *string, maxResults, startAt int) (output []byte, err error) {
	const path = "/rest/api/3/search"

	payload := JQL.SearchQuery{
		Expand:       []string{},
		Fields:       []string{},
		FieldsByKeys: false,
		Jql:          *query,
		MaxResults:   maxResults,
		StartAt:      startAt,
	}

	return jira.client.Send(
		http.MethodPost,
		path,
		payload.Bytes())
}
