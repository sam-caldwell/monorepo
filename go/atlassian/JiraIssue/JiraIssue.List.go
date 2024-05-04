package JiraIssue

import (
	"bytes"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
)

// List - list issues given a JQL Query object
func (jira *Issue) List(domain *Atlassian.Domain, jql *AtlassianTypes.JqlQuery) (*http.Request, error) {
	const path = "/rest/api/2/issue/"

	return http.NewRequest(
		http.MethodPost,
		Atlassian.JiraUrlFactory(Atlassian.JiraUrlPattern, domain.Get(), path),
		bytes.NewBuffer(jql.Bytes()))

}
