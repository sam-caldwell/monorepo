package JiraIssue

import (
	"github.com/sam-caldwell/monorepo/go/atlassian/JQL"
	"net/http"
)

// List - list issues given a JQL Query object
func (jira *Issue) List(query *string, expand, fields *[]string, maxResults, startAt uint) (output []byte, err error) {
	const path = "/rest/api/3/search"

	payload := JQL.SearchQuery{
		Expand:       *expand,
		Fields:       *fields,
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
