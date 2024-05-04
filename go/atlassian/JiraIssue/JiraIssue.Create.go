package JiraIssue

import (
	"bytes"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create(domain *Atlassian.Domain) (*http.Request, error) {
	const path = "/rest/api/2/issue/"

	return http.NewRequest(
		http.MethodPost,
		Atlassian.JiraUrlFactory(Atlassian.JiraUrlPattern, domain.Get(), path),
		bytes.NewBuffer(jira.Marshall()))

}
